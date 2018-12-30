package main

import (
	"context"
	"fmt"
	"log"

	"github.com/FernandoDevBh/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

const port = "50051"
const address = "0.0.0.0"

func main() {
	cc, err := grpc.Dial(fmt.Sprintf("%v:%v", address, port), grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	fmt.Printf("Request Calculator from %v:%v\n", address, port)

	doCalculator(c)
}

func doCalculator(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &calculatorpb.SumRequest{
		Sum: &calculatorpb.Sum{
			FirstNumber:  3,
			SecondNumber: 10,
		},
	}

	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Calculator RPC: %v\n", err)
	}

	log.Printf("Response from Calculator: %v\n", res.Result)
}
