package main

import (
	pb "github.com/golovers/grpc/v2/api-golang/gr/greeting/v1"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"encoding/json"
)

func main() {
	// GRPC
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
	log.Println("Result from GRPC", rs)

	// REST
	rsp, err := http.Get("http://localhost:8080/v1/greeting")
	if err != nil {
		log.Fatal(err)
	}
	rspJson := &pb.GreetingResponse{}
	json.NewDecoder(rsp.Body).Decode(rspJson)
	log.Println("Result from REST", rspJson)

}
