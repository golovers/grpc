package main

import (
	"google.golang.org/grpc"
	"net"
	"log"
	"github.com/golovers/grpc/v1/server/service"
	pb "github.com/golovers/grpc/v1/proto"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("Failed to start greeting", err)
	}
	server := grpc.NewServer()
	greetings := service.NewGreetingService()
	pb.RegisterGreetingsServer(server, greetings)

	server.Serve(lis)
}
