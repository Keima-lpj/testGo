package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	hello "testGo/grpc/proto"
)

type HelloServer struct {
}

func (h *HelloServer) HelloWorld(ctx context.Context, req *hello.Request) (*hello.Response, error) {
	return &hello.Response{Res: "This is the server " + req.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	hello.RegisterHelloServer(grpcServer, &HelloServer{})
	grpcServer.Serve(listen)
}
