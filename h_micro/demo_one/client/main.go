package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	proto "higo/h_micro/demo_one/proto"
)

func main() {
	//创建一个服务，并配置一些选项
	service := micro.NewService(
		micro.Name("greeter.client"),
	)

	//服务初始化
	service.Init()

	//创建一个客户端
	greeter := proto.NewGreeterService("greeter", service.Client())

	//调用这个Greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	//打印结果
	fmt.Println(rsp.Greeting)
}
