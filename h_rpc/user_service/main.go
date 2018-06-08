package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"higo/h_rpc/user_service/handle"
	proto "higo/h_rpc/user_service/proto"
)

const Topic = "user.created"

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	//获取broker实例
	//pubSub := srv.Server().Options().Broker
	publisher := micro.NewPublisher("user.created", srv.Client())

	token := handle.TokenService{}
	handler := &handle.Handler{TokenService: &token, Publisher: publisher}
	proto.RegisterUserServiceHandler(srv.Server(), handler)

	if err := srv.Run(); err != nil {
		fmt.Printf("user service error : %v\n", err)
	}
}
