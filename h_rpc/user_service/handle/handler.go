package handle

import (
	"context"
	"errors"
	"fmt"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/nats"
	proto "higo/h_rpc/user_service/proto"
)

type Handler struct {
	TokenService AuthAbel
	Publisher    micro.Publisher
	//PubSub       broker.Broker
}

// 发送消息通知
/*func (h *Handler) publishEvent(user *proto.User) error {
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	msg := &broker.Message{
		Header: map[string]string{
			"ip": user.Ip,
		},
		Body: body,
	}

	// 发布 user.created topic 消息
	if err := h.PubSub.Publish("user.created", msg); err != nil {
		log.Fatalf("[pub] failed: %v\n", err)
	}
	return nil
}*/

func (h *Handler) Create(ctx context.Context, req *proto.User, rsp *proto.Response) error {
	fmt.Println("2121212")
	// 发布带有用户所有信息的消息
	/*if err := h.publishEvent(req); err != nil {
		return err
	}*/

	if err := h.Publisher.Publish(ctx, req); err != nil {
		return err
	}

	fmt.Printf("Create : %v\n", req.Ip)
	return nil
}

func (h *Handler) Get(ctx context.Context, req *proto.User, rsp *proto.Response) error {
	fmt.Println("dsdsdsd", req.Ip)
	/*if err := h.Publisher.Publish(ctx, req); err != nil {
		return err
	}*/
	u := &proto.User{
		Ip: "33333",
	}
	rsp.User = u
	return nil
}

func (h *Handler) GetAll(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	return nil
}

func (h *Handler) Auth(ctx context.Context, req *proto.User, rsp *proto.Token) error {
	/*u := &proto.User{
		Password: "12345",
	}

	//验证
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return err
	}*/

	u := &proto.User{
		Ip:       req.Ip,
		Password: req.Password,
	}
	//生成token
	token, err := h.TokenService.Encode(u)
	if err != nil {
		return err
	}

	rsp.Token = token
	return nil
}

func (h *Handler) ValidateToken(ctx context.Context, req *proto.Token, rsp *proto.Token) error {
	claims, err := h.TokenService.Decode(req.Token)
	if err != nil {
		return err
	}
	if claims.Ip == "" {
		return errors.New("invalid user")
	}

	rsp.Valid = true
	return nil
}
