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
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)

	srv.Init()

	//注册订阅服务
	micro.RegisterSubscriber(Topic, srv.Server(), new(Subscriber))

	//运行服务
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

func (sub *Subscriber) Process(ctx context.Context, user *proto.RegisterRequest) error {
	log.Println("[Sending email to] : ", user.Email)
	return nil
}
