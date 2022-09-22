package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"sample.com/grpc-server-with-go/internal/service"
	"sample.com/grpc-server-with-go/pb"
)

func main() {
	port := 3000
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterSampleServiceServer(server, service.NewSampleServiceServer())
	reflection.Register(server)
	server.Serve(listenPort)
}
