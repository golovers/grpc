package main

import (
	pb "github.com/golovers/grpc/v1/proto"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to daial to greeting greeting", err)
	}
	greetings := pb.NewGreetingsClient(conn)
	rs, err := greetings.Say(context.Background(), &pb.GreetingRequest{
		Name: "Jack",
		Type: "HPBD",
	})
	if err != nil {
		log.Fatal("Fail to call to greeting greeting", err)
	}
	log.Println(rs)
}
