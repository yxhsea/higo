package main

import (
	"github.com/micro/go-micro"
	greeter "higo/h_micro/demo_six/srv/proto/greeter"
	"log"
)

type Say struct {
	Client greeter.SayService
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.greeter"),
	)

	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(&Say{Client: greeter.NewSayService("go.micro.srv.greeter", service.Client())}),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
		return
	}
}
