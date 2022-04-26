package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/nXnUs25/gRPC-factorial/cli"
	"github.com/nXnUs25/gRPC-factorial/rpc_factorial/proto"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmsgprefix)
	log.SetPrefix("FactorialReciver: ")
	log.SetOutput(os.Stdout)

	values := cli.Cmds()
	log.Printf("Starting... GRPC Client to talk with GRPC Factorial Server on port %d\n", cli.Port)

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", "localhost", cli.Port), grpc.WithInsecure())
	if err != nil {
		log.Printf("")
		log.Fatal(err)
	}
	defer conn.Close()
	client := proto.NewFactorialClient(conn)

	talking(client, values)
}

func talking(pfc proto.FactorialClient, nums []int64) {
	req := &proto.CalculateRequest{
		Numbers: nums,
	}
	process, err := pfc.Calculate(context.Background(), req)
	if err != nil {
		log.Println("Error processing calation requests")
		log.Fatal(err)
	}

	for _, num := range nums {
		msg, err := process.Recv()
		if err != nil {
			log.Printf("Error processing request for %v", num)
			log.Fatal(err)
		}
		fmt.Printf("The factorial of [%v] is [%v]\n", msg.GetInputNumber(), msg.GetFactorialResult())
	}
}
