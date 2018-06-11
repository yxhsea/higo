package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"
	greeter "higo/h_micro/demo_six/srv/proto/greeter"
	"log"
)

type Say struct {
}

var cl greeter.SayService

func (s *Say) Anything(c *gin.Context) {
	log.Print("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API.",
	})
}

func (s *Say) Hello(c *gin.Context) {
	log.Print("Received Say.Hello API request")
	name := c.Param("name")

	response, err := cl.Hello(context.TODO(), &greeter.Request{Name: name})
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, response)
}

func main() {
	service := web.NewService(
		web.Name("go.micro.api.greeter"),
	)

	service.Init()

	cl = greeter.NewSayService("go.micro.srv.greeter", client.DefaultClient)

	say := new(Say)
	router := gin.Default()
	router.GET("/greeter", say.Anything)
	router.GET("/greeter/:name", say.Hello)

	service.Handle("/", router)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
