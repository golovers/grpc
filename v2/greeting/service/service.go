package service

import (
	"fmt"
	pb "github.com/golovers/grpc/v2/api-golang/gr/greeting/v1"
	"golang.org/x/net/context"
)

type Greetings struct{}

func NewGreetingService() *Greetings {
	return &Greetings{}
}

func (gr *Greetings) Say(ctx context.Context, rq *pb.GreetingRequest) (*pb.GreetingResponse, error) {

	if rq.GetType() == "HPBD" {
		return &pb.GreetingResponse{
			Message: &pb.Greeting{
				Title:   fmt.Sprintf("Happy birthday to %v", rq.GetName()),
				Message: "We wish you a happy birthday",
			},
		}, nil
	}
	return &pb.GreetingResponse{
		Message: &pb.Greeting{
			Title:   fmt.Sprintf("Hello %v", rq.GetName()),
			Message: "We wish you a happy day",
		},
	}, nil
}
