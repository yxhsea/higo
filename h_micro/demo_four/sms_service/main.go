package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	proto "higo/h_micro/demo_four/user_service/proto"
	"log"
)

const Topic = "user.register"

type Subscriber struct {
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.sms"),
		micro.Version("latest"),
	)

	micro.RegisterSubscriber(Topic, srv.Server(), new(Subscriber))

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

func (sub *Subscriber) Process(ctx context.Context, user *proto.RegisterRequest) error {
	log.Println("[Sending Sms to] : ", user.Phone)
	return nil
}
