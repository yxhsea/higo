package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	proto "higo/h_micro/demo_three/consignment_service/proto"
	userPb "higo/h_micro/demo_three/user_service/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)

	//解析命令行参数
	service.Init()

	//注册
	proto.RegisterShippingServiceHandler(service.Server(), &Handler{})

	//运行
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		token := meta["Token"]

		authClient := userPb.NewUserService("go.micro.srv.user", client.DefaultClient)
		authResp, err := authClient.ValidateToken(context.Background(), &userPb.Token{
			Token: token,
		})
		fmt.Println("Auth Resp: ", authResp)

		if err != nil {
			return err
		}
		err = fn(ctx, req, rsp)
		return nil
	}
}
