package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	//"github.com/micro/go-micro/metadata"
	proto "higo/h_rpc/hello_service/proto"
	"os"
	"time"
)

func main() {
	//创建微服务客户端
	cli := proto.NewHelloService("go.micro.srv.hello", client.DefaultClient)

	/*token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpcCI6ImYwZmRiNGMzZjU4ZTNlM2Y4ZTc3MTYyZDg5M2QzMDU1IiwicGFzc3dvcmQiOiJjYzAzZTc0N2E2YWZiYmNiZjhiZTc2NjhhY2ZlYmVlNSIsImV4cCI6MTUyODYxOTI2MywiaXNzIjoiZ28ubWljcm8uc3J2LnVzZXIifQ.sJrfSawdl1JnE9AuHJxD01FpHL2hSAcTs_ghPH3WGIM"
	//token := ""
	tokenContext := metadata.NewContext(context.TODO(), map[string]string{
		"token": token,
	})*/

	for i := 0; i < 10; i++ {
		go func() {
			rsp, err := cli.GetOne(context.TODO(), &proto.HelloRequest{Msg: "message"})
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			fmt.Println(rsp.Reply)
		}()
	}
	time.Sleep(10 * time.Minute)
}
