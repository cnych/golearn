package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	pb "scorpio/grpcapp/helloworld/helloworld"
	"starjazz/logx"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

//SayHello(context.Context, *HelloRequest) (*HelloReply, error)
func (*server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (*server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello Again " + in.GetName()}, nil
}

func main() {
	lisen, err := net.Listen("tcp", port)
	if err != nil {
		logx.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lisen); err != nil {
		logx.Fatalf("Failed to serve: %v", err)
	}
}
