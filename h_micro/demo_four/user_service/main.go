package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	proto "higo/h_micro/demo_four/user_service/proto"
)

const Topic = "user.register"

type User struct {
	Publisher micro.Publisher
}

func (u *User) Register(context context.Context, req *proto.RegisterRequest, rps *proto.RegisterResponse) error {
	rps.Id = "123456"
	fmt.Printf("[Register user information] emial : %v; phone: %v; password : %v\n", req.Email, req.Phone, req.Password)

	if err := u.Publisher.Publish(context, req); err != nil {
		fmt.Println(err)
	}
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()
	publisher := micro.NewPublisher(Topic, srv.Client())

	//注册服务
	proto.RegisterUserServiceHandler(srv.Server(), &User{publisher})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
