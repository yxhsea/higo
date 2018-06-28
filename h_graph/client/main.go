package main

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"time"
)

func main() {
	req := gorequest.New()
	rsp, body, err := req.Post("http://127.0.0.1:8080/query").Timeout(30 * time.Second).Send("{\"query\": \"{ hello }\"}").End()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp)
	fmt.Println(string(body))
}
