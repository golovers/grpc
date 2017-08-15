package main

import (
	pb "github.com/golovers/grpc/v3/api-golang/gr/greeting/v1"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"encoding/json"
	"flag"
)

var
(
	HTTP_ENDPOINT = flag.String("http endpoint", "http://192.168.99.100:30100", "")
	GRPC_ENDPOINT = flag.String("grpc endpoint", "192.168.99.100:30101", "")
)

func main() {
	// GRPC
	conn, err := grpc.Dial(*GRPC_ENDPOINT, grpc.WithInsecure())
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
	log.Println("Result from GRPC",rs)

	// REST
	rsp,err := http.Get(*HTTP_ENDPOINT +"/v1/greeting")
	if err != nil {
		log.Fatal(err)
	}
	rspJson := &pb.GreetingResponse{}
	json.NewDecoder(rsp.Body).Decode(rspJson)
	log.Println("Result from REST", rspJson)

}
