package main

import "context"
import (
	"github.com/micro/go-micro"
	greeter "higo/h_micro/demo_six/srv/proto/greeter"
	"log"
	"time"
)

type Say struct {
}

func (s *Say) Hello(ctx context.Context, req *greeter.Request, rsp *greeter.Response) error {
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()

	greeter.RegisterSayHandler(service.Server(), new(Say))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
