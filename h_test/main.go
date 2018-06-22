package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	proto "higo/h_test/proto"
)

func main() {
	service := micro.NewService(
		micro.Registry(consul.NewRegistry(consul.Config(&api.Config{Address: "192.168.0.163:8500"}))),
	)
	service.Init()

	//创建user_service微服务客户端
	cli := proto.NewGenerateUniqueIdService("go.micro.srv.unique_id", service.Client())
	rsp, err := cli.GenerateUniqueId(context.TODO(), &proto.UniqueIdRequest{Key: "testId"})
	if err != nil {
		fmt.Printf("Generate error: %v", err)
	} else {
		fmt.Println("Generate id is: ", rsp.Value)
	}
}
