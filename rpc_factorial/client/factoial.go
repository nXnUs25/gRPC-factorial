package main

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
)

var gRPC_Port = 5100

func main() {
	log.SetFlags(log.LstdFlags | log.Lmsgprefix)
	log.SetPrefix("FactorialReciver: ")
	log.SetOutput(os.Stdout)
	SetGRPCPort(GRPCPort())
	log.Printf("Starting... GRPC server on port %d\n", GRPCPort())

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", "localhost", GRPCPort()), grpc.WithInsecure())
	if err != nil {
		log.Printf("")
		log.Fatal(err)
	}
	defer conn.Close()

}

func GRPCPort() int {
	return gRPC_Port
}

func SetGRPCPort(port int) {
	gRPC_Port = port
}
