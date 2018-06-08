package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	proto "higo/h_micro/demo_five/user_service/proto"
)

func main() {
	cmd.Init()
	//创建user_service微服务客户端
	cli := proto.NewCreateUserService("go.micro.srv.user", client.DefaultClient)
	rsp, err := cli.CreateUser(context.TODO(), &proto.UserRequest{Username: "yxhsea", Password: "yxhsea"})
	if err != nil {
		fmt.Printf("call create error: %v", err)
	} else {
		fmt.Println("craeted: ", rsp)
	}
}
