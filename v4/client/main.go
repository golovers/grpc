package main

import (
	pb "github.com/golovers/grpc/v4/api-golang/gr/greeting/v1"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"encoding/json"
	"flag"
	"google.golang.org/grpc/credentials"
	"crypto/tls"
)

var
(
	HTTP_ENDPOINT = flag.String("http endpoint", "https://192.168.99.100:30100", "")
	GRPC_ENDPOINT = flag.String("grpc endpoint", "192.168.99.100:30101", "")
	TLSServerName = "localhost"
)

func main() {
	// GRPC8
	cred := credentials.NewClientTLSFromCert(nil, TLSServerName)
	conn, err := grpc.Dial(*GRPC_ENDPOINT, grpc.WithTransportCredentials(cred))
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
	mTLSConfig := &tls.Config {
		ServerName:"localhost",
	}

	tr := &http.Transport{
		TLSClientConfig: mTLSConfig,
	}

	client := &http.Client{Transport: tr}
	rsp,err1 := client.Get(*HTTP_ENDPOINT +"/v1/greeting")

	if err1 != nil {
		log.Fatal(err1)
	}

	rspJson := &pb.GreetingResponse{}
	json.NewDecoder(rsp.Body).Decode(rspJson)
	log.Println("Result from REST", rspJson)

}
