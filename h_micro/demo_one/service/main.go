package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	proto "higo/h_micro/demo_one/proto"
)

type Greeter struct {
}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	//创建一个服务，并配置一些选项
	service := micro.NewService(
		micro.Name("greeter"),
	)

	//解析命令行参数
	service.Init()

	greeter := &Greeter{}
	//注册服务处理
	proto.RegisterGreeterHandler(service.Server(), greeter)

	//运行这个服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
