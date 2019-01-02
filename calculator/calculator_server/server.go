package main

import (
	"context"
	"fmt"
	"io"
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

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("PrimeNumberDecomposition function was invoked with %v\n", req)
	div := int64(2)
	number := req.GetPrimeDecompositon().GetNumber()
	for number > 1 {
		if number%div == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: div,
			})
			number = number / div
		} else {
			div += 1
		}
	}
	return nil
}

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("ComputeAverage function was invoked with streaming request: %v\n", stream)
	var counter float64
	var sum float64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// We have finished reading the client stream
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Average: sum / counter,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		counter++
		sum += req.AverageComposition.GetNumber()
	}
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
