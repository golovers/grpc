package main

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/golovers/grpc/v2/api-golang/gr/greeting/v1"
	"github.com/golovers/grpc/v2/greeting/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	// gRPC
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("Failed to start greeting", err)
	}
	server := grpc.NewServer()
	greetings := service.NewGreetingService()
	pb.RegisterGreetingsServer(server, greetings)

	// HTTP
	mux := runtime.NewServeMux()
	dopts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	if err := pb.RegisterGreetingsHandlerFromEndpoint(context.Background(), mux, ":8000", dopts); err != nil {
		log.Fatal(err)
	}

	go server.Serve(lis)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
