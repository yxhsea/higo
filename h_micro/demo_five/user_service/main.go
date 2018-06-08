package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	proto "higo/h_micro/demo_five/user_service/proto"
)

type User struct {
}

func (u *User) CreateUser(context context.Context, req *proto.UserRequest, rsp *proto.UserResponse) error {
	fmt.Println("user information : ", req)
	rsp.Id = "123456"
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	service.Init()
	proto.RegisterCreateUserHandler(service.Server(), new(User))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
