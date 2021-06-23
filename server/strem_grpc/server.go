package main

import (
	"gRPC/pkg"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/EDDYCJY/go-grpc-example/proto"

)



const (
	PORT = "9002"
)

func main() {
	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server, &pkg.StreamService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}

