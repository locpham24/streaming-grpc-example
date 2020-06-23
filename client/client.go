package main

import (
	"context"
	"fmt"
	pb "github.com/locpham24/streaming-grpc-example/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func runUnaryRPC(client pb.NumberClient, input *pb.InputNumber) error {
	res, err := client.HelloNumber(context.Background(), input)
	if err != nil {
		return err
	}
	fmt.Println(res.Greeting)
	return nil
}

func runServerStreamingRPC(client pb.NumberClient, input *pb.InputNumber) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	stream, err := client.MultipleOfTwo(ctx, input)
	if err != nil {
		log.Fatalf("%v.MultipleOfTwo(_) = _, %v", client, err)
	}

	for {
		out, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Finish")
			break
		}
		if err != nil {
			log.Fatalf("%v.MultipleOfTwo(_) = _, %v", client, err)
		}
		fmt.Println(out.Num)
	}
	return nil
}

func main() {
	conn, _ := grpc.Dial("localhost:50001", grpc.WithInsecure())

	client := pb.NewNumberClient(conn)

	input := &pb.InputNumber{
		Num: 2,
	}

	//runUnaryRPC(client, input)
	runServerStreamingRPC(client, input)

}
