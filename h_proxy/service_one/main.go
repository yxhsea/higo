package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/abc", func(context *gin.Context) {
		fmt.Println("This is service one testing...")
	})

	err := r.Run(":8081")
	if err != nil {
		log.Fatal(err)
	}
}
