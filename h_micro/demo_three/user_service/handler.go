package user_service

import "context"
import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	proto "higo/h_micro/demo_three/user_service/proto"
)

type Handler struct {
	tokenService AuthAbel
}

func (h *Handler) Create(ctx context.Context, req *proto.User, rsp *proto.Response) error {
	return nil
}

func (h *Handler) Get(ctx context.Context, req *proto.User, rsp *proto.Response) error {
	return nil
}

func (h *Handler) GetAll(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	return nil
}

func (h *Handler) Auth(ctx context.Context, req *proto.User, rsp *proto.Token) error {
	u := &proto.User{
		Password: "12345",
	}

	//验证
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return err
	}

	//生成token
	token, err := h.tokenService.Encode(u)
	if err != nil {
		return err
	}

	rsp.Token = token
	return nil
}

func (h *Handler) ValidateToken(ctx context.Context, req *proto.Token, rsp *proto.Token) error {
	claims, err := h.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}
	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	rsp.Valid = true
	return nil
}
