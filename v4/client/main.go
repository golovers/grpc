package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	pb "github.com/golovers/grpc/v4/api-golang/gr/greeting/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"net/http"
	"strconv"
)

var (
	HTTP_ENDPOINT = flag.String("http endpoint", "https://192.168.99.100:30100", "")
	GRPC_ENDPOINT = flag.String("grpc endpoint", "192.168.99.100:30101", "")
	TLSServerName = "localhost"
	JWT_TOKEN     = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJwdGhldGhhbmgiLCJleHAiOjE1MDc2MzY4MDAsImlhdCI6MTUwMzAzODA5MCwiaXNzIjoiZ29sb3ZlcnMiLCJuYmYiOjE1MDMwMzgwOTAsInN1YiI6IkpXVCJ9.GoL1oqCFkH_O9WZxaSHLMq57GkEo4wof655YAKPwlOk"
)

func main() {
	// GRPC8
	cred := credentials.NewClientTLSFromCert(nil, TLSServerName)
	conn, err := grpc.Dial(*GRPC_ENDPOINT, grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatal("Failed to daial to greeting greeting", err)
	}
	greetings := pb.NewGreetingsClient(conn)

	// Pass JWT for authentication
	md := metadata.Pairs("authorization", JWT_TOKEN)
	ctx := context.Background()
	ctx = metadata.NewContext(ctx, md)
	rs, err := greetings.Say(ctx, &pb.GreetingRequest{
		Name: "Jack",
		Type: "HPBD",
	})
	if err != nil {
		log.Fatal("Fail to call to greeting greeting", err)
	}
	log.Println("Result from GRPC", rs)

	// Streaming
	maxMsg := 1000000000
	stream, err := greetings.SayAlots(ctx, []grpc.CallOption{
		grpc.MaxCallRecvMsgSize(maxMsg),
		grpc.MaxCallSendMsgSize(maxMsg),
	}...)
	// start the receiver
	go func() {
		for {
			if resp, err := stream.Recv(); err != nil {
				if err == io.EOF {
					log.Println("Finished receiving from server")
					break
				}
				log.Println("Got error", err)
			} else {
				log.Println("Received", resp)
			}
		}
	}()
	// start the sender
	go func() {
		for i := 0; i < 1000; i++ {
			log.Println("Send ", strconv.Itoa(i))
			stream.Send(&pb.GreetingRequest{
				Name: "Jack " + strconv.Itoa(i),
				Type: "HPBD",
			})
		}
		stream.CloseSend()
	}()

	// REST
	mTLSConfig := &tls.Config{
		ServerName: "localhost",
	}

	tr := &http.Transport{
		TLSClientConfig: mTLSConfig,
	}

	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", *HTTP_ENDPOINT+"/v1/greeting", nil)
	req.Header.Set("authorization", JWT_TOKEN)

	rsp, err1 := client.Do(req)
	if err1 != nil || rsp.StatusCode != http.StatusOK {
		log.Fatal(err1, rsp.Status)
	}

	rspJson := &pb.GreetingResponse{}
	json.NewDecoder(rsp.Body).Decode(rspJson)
	log.Println("Result from REST", rspJson)

	// Wait
	ch := make(chan int)
	<-ch
}
