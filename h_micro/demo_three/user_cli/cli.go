package main

import "github.com/micro/go-micro/cmd"
import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	proto "higo/h_micro/demo_three/user_service/proto"
	"os"
)

func main() {
	cmd.Init()

	//创建user_service微服务客户端
	cli := proto.NewUserService("go.micro.srv.user", client.DefaultClient)

	//暂时写死用户信息
	name := "Ewan Valentine"
	email := "ewan.valentine89@gmail.com"
	password := "test123"
	company := "BBC"

	rsp, err := cli.Create(context.TODO(), &proto.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		fmt.Printf("call create error: %v", err)
	}
	fmt.Println("craeted: ", rsp.User.Id)

	allRsp, err := cli.GetAll(context.Background(), &proto.Request{})
	if err != nil {
		fmt.Printf("call GetAll error : %v", err)
	}
	for i, u := range allRsp.Users {
		fmt.Printf("user_%d, %v\n", i, u.Id)
	}

	authRsp, err := cli.Auth(context.TODO(), &proto.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		fmt.Printf("auth failed: %v", err)
	}
	fmt.Println("token: ", authRsp.Token)

	//直接退出
	os.Exit(0)
}
