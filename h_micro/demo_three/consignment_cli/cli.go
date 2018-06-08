package main

import "github.com/micro/go-micro/cmd"
import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	proto "higo/h_micro/demo_three/consignment_service/proto"
	"os"
)

func main() {
	cmd.Init()

	//创建微服务客户端
	cli := proto.NewShippingService("go.micro.srv.consignment", client.DefaultClient)

	//在命令行中指定新的货物信息json件
	if len(os.Args) < 3 {
		fmt.Println("Not enough argments, expecting file and tokan.")
	}
	//	infoFile := os.Args[1]
	token := os.Args[2]

	//解析货物信息

	//创建带有用户token的context
	//consignment-service 服务端将从中取出token, 解密取出用户身份
	tokenContext := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	//调用Rpc
	//将货物存储到指定用户的仓库里
	rsp, err := cli.CreateConsignment(tokenContext, &proto.Consignment{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("created : %t", rsp.Created)

	for k, v := range rsp.Consignments {
		fmt.Printf("consignment_%d: %v\n", k, v)
	}
}
