package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/FernandoDevBh/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

const port = "50051"
const address = "0.0.0.0"
const network = "tcp"

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Calculator function was invoked with %v\n", req)
	fn := req.GetSum().GetFirstNumber()
	sn := req.GetSum().GetSecondNumber()
	rs := &calculatorpb.SumResponse{
		Result: fn + sn,
	}

	return rs, nil
}

func main() {
	lis, err := net.Listen(network, fmt.Sprintf("%v:%v", address, port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	fmt.Printf("Serving on address %v:%v", address, port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Server: %v", err)
	}
}
