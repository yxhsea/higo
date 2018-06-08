package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	proto "higo/h_micro/demo_four/user_service/proto"
	"time"
)

func main() {
	cli := proto.NewUserService("go.micro.srv.user", client.DefaultClient)

	fmt.Println("==========start===========")
	for i := 0; i < 50; i++ {
		go func() {
			req := &proto.RegisterRequest{
				Email:    "yxhsea@foxmail.com",
				Phone:    "18373288691",
				Password: "123456",
			}
			rsp, err := cli.Register(context.TODO(), req)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(rsp)
		}()
	}
	//fmt.Println("user_id", rsp.Id)
	time.Sleep(10 * time.Minute)
}
