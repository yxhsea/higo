package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
)

type Post struct {
	Title    string `valid:"alphanum,required"`
	Message  string `valid:"ascii"`
	AuthorIP string `valid:"ipv4"`
	Date     string `valid:"-"`
	Phone    string `valid:"numeric,required"`
}

func main() {
	post := &Post{
		Title:    "MyExamplePost",
		Message:  "duck",
		AuthorIP: "123.234.54.3",
		Phone:    "13838381238",
	}

	// Add your own struct validation tags
	govalidator.TagMap["duck"] = govalidator.Validator(func(str string) bool {
		return str == "duck"
	})

	result, err := govalidator.ValidateStruct(post)
	if err != nil {
		fmt.Println("请提交正确参数")
		println("error: " + err.Error())
	}
	println(result)
}
