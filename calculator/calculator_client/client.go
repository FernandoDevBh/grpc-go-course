package main

import (
	"context"
	"fmt"
	"io"
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

	//doCalculatorSum(c)

	doCalculatorPrimeDecom(c)
}

func doCalculatorSum(c calculatorpb.CalculatorServiceClient) {
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

func doCalculatorPrimeDecom(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Streaming RPC...")
	resStream, err := c.PrimeNumberDecomposition(context.Background(), &calculatorpb.PrimeNumberDecompositionRequest{
		PrimeDecompositon: &calculatorpb.PrimeDecompositon{
			Number: 120,
		},
	})

	if err != nil {
		log.Fatalf("error while calling PrimeNumberDecomposition RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from PrimeNumberDecomposition: %v\n", msg.GetPrimeFactor())
	}
}
