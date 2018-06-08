package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	proto "higo/h_rpc/hello_service/proto"
	userPb "higo/h_rpc/user_service/proto"
	"log"
)

const Topic = "user.created"

type Subscriber struct{}

type Hello struct {
}

func (h *Hello) GetOne(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	fmt.Printf("Message : %v\n", req.Msg)
	rsp.Reply = "reply"
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.hello"),
		micro.Version("latest"),
		//micro.WrapHandler(AuthWrapper),
	)

	//解析命令行参数
	service.Init()

	/*pubSub := service.Server().Options().Broker
	if err := pubSub.Connect(); err != nil {
		log.Fatalf("broker connect error: %v\n", err)
	}

	// 订阅消息
	_, err := pubSub.Subscribe(topic, func(pub broker.Publication) error {
		var user *userPb.User
		if err := json.Unmarshal(pub.Message().Body, &user); err != nil {
			return err
		}
		log.Printf("[CREATE USER]: %v\n", user)
		go senEmail(user)
		return nil
	})
	if err != nil {
		log.Printf("sub error: %v\n", err)
	}*/

	//micro.RegisterSubscriber("user.created", service.Server(), new(Subscriber))

	//注册
	proto.RegisterHelloHandler(service.Server(), new(Hello))

	//运行
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func (sub *Subscriber) Process(ctx context.Context, user *userPb.User) error {
	log.Println("[Picked up a new message]")
	log.Println("[Sending email to]:", user.Ip)
	return nil
}

//授权中间件
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}
		token := meta["Token"]
		if len(token) <= 0 {
			return errors.New("token is empty")
		}

		authClient := userPb.NewUserService("go.micro.srv.user", client.DefaultClient)
		authResp, err := authClient.ValidateToken(context.Background(), &userPb.Token{
			Token: token,
		})
		if err != nil || !authResp.Valid {
			return errors.New("auth failure")
		}

		err = fn(ctx, req, rsp)
		return nil
	}
}

func senEmail(user *userPb.User) error {
	log.Printf("[SENDING A EMAIL TO %s...]", user.Ip)
	return nil
}
