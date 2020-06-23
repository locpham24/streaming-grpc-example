package main

import (
	"context"
	"fmt"
	pb "github.com/locpham24/streaming-grpc-example/proto"
	"google.golang.org/grpc"
	"io"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type numberServer struct{}

func (n *numberServer) HelloNumber(ctx context.Context, in *pb.InputNumber) (*pb.HelloOutput, error) {
	return &pb.HelloOutput{
		Greeting: "Hello number: " + fmt.Sprintf("%d", in.Num),
	}, nil
}

func (n *numberServer) MultipleOfTwo(in *pb.InputNumber, stream pb.Number_MultipleOfTwoServer) error {
	for i := 1; i < 10000; i++ {
		out := &pb.OutputNumber{
			Num: in.Num * int64(i),
		}

		if err := stream.Send(out); err != nil {
			return err
		}
	}
	return nil
}

func (n *numberServer) SumOfNumbers(stream pb.Number_SumOfNumbersServer) error {
	var sum int64
	for {
		input, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.OutputNumber{
				Num: sum,
			})
		}
		if err != nil {
			return err
		}

		sum += input.Num
	}
	return nil
}

func (n *numberServer) NumberChat(stream pb.Number_NumberChatServer) error {
	for {
		input, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		output := &pb.OutputNumber{
			Num: input.Num * 10,
		}
		if err := stream.Send(output); err != nil {
			return err
		}
		delay := rand.Intn(1500) + 50
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Printf("%s: %d\n", strconv.Itoa(delay)+"ms", output.Num)
	}
	return nil
}

func main() {
	listener, _ := net.Listen("tcp", ":50001")
	grpcServer := grpc.NewServer()

	pb.RegisterNumberServer(grpcServer, &numberServer{})

	fmt.Println("Starting server at port 50001...")
	grpcServer.Serve(listener)
}
