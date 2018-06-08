package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	proto "higo/h_rpc/user_service/proto"
	"os"
	"time"
)

func main() {
	cmd.Init()

	//创建user_service微服务客户端
	cli := proto.NewUserService("go.micro.srv.user", client.DefaultClient)

	//暂时写死用户信息
	ip := "192.168.0.2"
	password := "test123"

	for i := 0; i < 10; i++ {
		//go func() {
		rsp, err := cli.Get(context.TODO(), &proto.User{Ip: ip, Password: password})
		if err != nil {
			fmt.Printf("call create error: %v", err)
			os.Exit(0)
		}
		fmt.Println("craeted: ", rsp.User)
		time.Sleep(time.Second)
		//}()
	}
	os.Exit(0)
	/*allRsp, err := cli.GetAll(context.Background(), &proto.Request{})
	if err != nil {
		fmt.Printf("call GetAll error : %v", err)
	}
	for i, u := range allRsp.Users {
		fmt.Printf("user_%d, %v\n", i, u.Ip)
	}*/

	authRsp, err := cli.Auth(context.TODO(), &proto.User{
		Ip:       ip,
		Password: password,
	})
	if err != nil {
		fmt.Printf("auth failed: %v", err)
		os.Exit(0)
	}
	fmt.Println("token: ", authRsp.Token)

	tokenRsp, err := cli.ValidateToken(context.TODO(), &proto.Token{Token: authRsp.Token})
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Printf("Token %v, vaild : %v", tokenRsp.Token, tokenRsp.Valid)

	//直接退出
	os.Exit(0)
}
