package main

import (
	"context"
	"fmt"
	pb "github.com/locpham24/streaming-grpc-example/proto"
	"google.golang.org/grpc"
	"net"
)

type numberServer struct{}

func (n *numberServer) HelloNumber(ctx context.Context, in *pb.InputNumber) (*pb.HelloOutput, error) {
	return &pb.HelloOutput{
		Greeting: "Hello number: " + fmt.Sprintf("%d", in.Num),
	}, nil
}

func main() {
	listener, _ := net.Listen("tcp", ":50001")
	grpcServer := grpc.NewServer()

	pb.RegisterNumberServer(grpcServer, &numberServer{})

	fmt.Println("Starting server at port 50001...")
	grpcServer.Serve(listener)
}
