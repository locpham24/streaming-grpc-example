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

func runClientStreamingRPC(client pb.NumberClient) error {
	var inputs []*pb.InputNumber
	for i := 1; i < 10; i++ {
		in := &pb.InputNumber{
			Num: int64(i),
		}
		inputs = append(inputs, in)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	stream, err := client.SumOfNumbers(ctx)
	if err != nil {
		log.Fatalf("%v.SumOfNumbers(_) = _, %v", client, err)
	}

	for _, input := range inputs {
		if err := stream.Send(input); err != nil {
			return err
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	fmt.Println(reply.Num)
	return nil
}

func runBidiStreamingRPC(client pb.NumberClient) error {
	var inputs []*pb.InputNumber
	for i := 1; i < 10000; i++ {
		in := &pb.InputNumber{
			Num: int64(i),
		}
		inputs = append(inputs, in)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stream, err := client.NumberChat(ctx)
	if err != nil {
		log.Fatalf("%v.NumberChat(_) = _, %v", client, err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			out, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("%v.NumberChat(_) = _, %v", client, err)
			}
			fmt.Println("receive: ", out.Num)
		}
	}()
	for _, input := range inputs {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("sent: ", input.Num)
		if err := stream.Send(input); err != nil {
			return err
		}
	}
	stream.CloseSend()
	<-waitc
	return nil
}
func main() {
	conn, _ := grpc.Dial("localhost:50001", grpc.WithInsecure())

	client := pb.NewNumberClient(conn)

	/*input := &pb.InputNumber{
		Num: 2,
	}*/

	//runUnaryRPC(client, input)
	//runServerStreamingRPC(client, input)
	//runClientStreamingRPC(client)
	runBidiStreamingRPC(client)

}
