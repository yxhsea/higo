package main

import "github.com/micro/go-micro"
import (
	"context"
	"fmt"
	greeter "higo/h_micro/demo_six/srv/proto/greeter"
	"log"
)

func main() {
	service := micro.NewService()

	service.Init()

	cl := greeter.NewSayService("go.micro.srv.greeter", service.Client())

	rsp, err := cl.Hello(context.Background(), &greeter.Request{Name: "John"})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(rsp.Msg)
}
