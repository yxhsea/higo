package main

import "context"
import proto "higo/h_micro/demo_three/consignment_service/proto"

type Handler struct {
}

func (h *Handler) CreateConsignment(ctx context.Context, req *proto.Consignment, rsp *proto.Response) error {
	return nil
}

func (h *Handler) GetConsignment(ctx context.Context, req *proto.GetRequest, rsp *proto.Response) error {
	return nil
}
