package main

import (
	"context"
	"google.golang.org/grpc"
	"os"
	pb "scorpio/grpcapp/helloworld/helloworld"
	"starjazz/logx"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logx.Fatalf("dia not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		logx.Fatalf("could not greet: %v", err)
	}
	logx.Printf("Greeting: %s", r.Message)

	r, err = c.SayHelloAgain(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		logx.Fatalf("could not greet: %v", err)
	}
	logx.Printf("Greeting: %s", r.Message)
}
