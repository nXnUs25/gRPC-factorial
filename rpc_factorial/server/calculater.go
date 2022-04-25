package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	f "github.com/nXnUs25/gRPC-fact/factorial"
	pb3 "github.com/nXnUs25/gRPC-factorial/rpc_factorial/proto"
)

const (
	varPort = "GRPC_PORT"
	varPrec = "FACTORIAL_PRECISION"
)

var (
	gRPC_Port int
	prec      uint
)

type gRPCServer struct {
	pb3.UnimplementedFactorialServer
}

func (r *gRPCServer) Calculate(req *pb3.CalculateRequest, cs *pb3.CalculateServer) error {
	nums := req.GetNumbers()
	var cal f.Calculater
	for _, num := range nums {
		if num > 0 {
			cal = f.MakeCalculate(num)
			result := cal.Calculate(cal, prec)
			cs.Send(&pb3.CalculateResult{
				InputNumber:     num,
				FactorialResult: result,
			})
			return nil
		}
		return status.Errorf(codes.InvalidArgument, "Invalid number [%v], Factorial can calcualte only positive numbers", num)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmsgprefix)
	log.SetPrefix("gRPCServerFactorial: ")
	log.SetOutput(os.Stdout)
	log.Printf("Starting... GRPC server on port %d\n", GRPCPort())

	l, err := grpc.Listen("tcp", "127.0.0.1:"+GRPCPort())
	if err != nil {
		log.Printf("Cannot open port [%v]", GRPCPort())
		log.Fatal(err)
	}
	server := grpc.NewServer()
	defer server.Stop()
	pb.RegisterFactorialServer(server, &gRPCServer{})
	log.Printf("Server started at port [%v]", l.Addr().String)
	if err := server.Serve(); err != nil {
		log.Println("Failed to Start")
		log.Fatal(err)
	}
	sChan := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sChan
		log.Printf("\nGot signal [%v]", sig)
		done <- true
	}()
	<-done
	log.Println("Sopping GRPC Server.")
	server.GracefulStop()
}

func GRPCPort() int {
	return gRPC_Port
}

func SetGRPCPort(port int) {
	gRPC_Port = port
}

func readGRPCPort() int {
	port, err := strconv.Atoi(os.Getenv(varPort))
	if err != nil {
		log.Printf("Failed to load port variable [%v] trying default port number [%v]", varPort, 5100)
		return 5100
	}
	return port
}

func Prec() uint {
	return prec
}

func readPrecision() uint {
	prec, err := strconv.Atoi(os.Getenv(varPrec))
	if err != nil {
		log.Printf("Failed to load precision variable [%v] trying default value number [%v]", varPrec, 64)
		return uint(64)
	}
	return uint(prec)
}
