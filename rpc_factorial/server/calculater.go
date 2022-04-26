package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"google.golang.org/grpc/status"

	"github.com/nXnUs25/gRPC-factorial/factorial"
	"github.com/nXnUs25/gRPC-factorial/rpc_factorial/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type grpcServer struct {
	proto.UnimplementedFactorialServer
}

type FactorialCaller interface {
	Calculate(int64) string
}

func (s *grpcServer) Calculate(req *proto.CalculateRequest, pfcs proto.Factorial_CalculateServer) error {
	nums := req.GetNumbers()

	var fact FactorialCaller = factorial.NewCounter()
	for _, num := range nums {
		if num < 0 {
			return status.Errorf(codes.InvalidArgument, "Negative number %+v. Only positive numbers allowed to calculate factorial", num)
		}
		result := fact.Calculate(num)
		pfcs.Send(&proto.CalculateResult{
			InputNumber:     num,
			FactorialResult: result,
		})
	}
	return nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmsgprefix)
	log.SetPrefix("gRPCServerFactorial: ")
	log.SetOutput(os.Stdout)
	SetGRPCPort(readGRPCPort())
	log.Printf("Starting... GRPC client on port %d\n", GRPCPort())

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", GRPCPort()))
	if err != nil {
		log.Printf("Cannot open port [%v]", GRPCPort())
		log.Fatal(err)
	}
	server := grpc.NewServer()
	defer server.Stop()
	proto.RegisterFactorialServer(server, &grpcServer{})
	log.Printf("Server started at port [%v]", listener.Addr())

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
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

const (
	varPort = "GRPC_PORT"
)

var (
	gRPC_Port int
)

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
