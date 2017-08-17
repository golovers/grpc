package service

import (
	"fmt"
	pb "github.com/golovers/grpc/v4/api-golang/gr/greeting/v1"
	"golang.org/x/net/context"
	"io"
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

func (gr *Greetings) SayAlots(stream pb.Greetings_SayAlotsServer) error {
	for {

		if req, err := stream.Recv(); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		} else {
			go func(req *pb.GreetingRequest, str pb.Greetings_SayAlotsServer) {
				if resp, err := gr.Say(context.Background(), req); err != nil {
					str.Send(&pb.GreetingResponse{})
				} else {
					str.Send(resp)
				}
			}(req, stream)
		}
	}
	return nil
}
