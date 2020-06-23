package main

import (
	"context"
	"fmt"
	pb "github.com/locpham24/streaming-grpc-example/proto"
	"google.golang.org/grpc"
)

func runUnaryRPC(client pb.NumberClient, input *pb.InputNumber) error {
	res, err := client.HelloNumber(context.Background(), input)
	if err != nil {
		return err
	}
	fmt.Println(res.Greeting)
	return nil
}

func main() {
	conn, _ := grpc.Dial("localhost:50001", grpc.WithInsecure())

	client := pb.NewNumberClient(conn)

	input := &pb.InputNumber{
		Num: 2,
	}

	runUnaryRPC(client, input)

}
