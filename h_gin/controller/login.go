package controller

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"higo/h_gin/pkg/util"
	"time"
)

type Login struct {
	UserName string `form:"user_name" json:"user_name" valid:"required,length(6|12)~用户名长度范围是6到12个字符"`
	PassWord string `form:"pass_word" json:"pass_word" valid:"required,length(6|12)~密码长度范围是6到12个字符"`
}

// @Summary 登录授权
// @tags auth
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param user_name formData string true "UserName"
// @Param pass_word formData string true "PassWord"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /login/auth [post]
func LoginAuth(ctx *gin.Context) {
	var login Login
	ctx.ShouldBind(&login)

	_, err := govalidator.ValidateStruct(login)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(401, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	}

	token, err := util.GenerateToken(login.UserName, login.PassWord)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(401, gin.H{
			"code": -1,
			"data": "",
			"msg":  "授权失败",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code": 0,
		"data": map[string]string{"token": token},
		"msg":  "授权成功",
	})
	return
}

type Token struct {
	Token string `form:"token" json:"token" valid:"required"`
}

// @Summary 登录授权检查
// @tags auth
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param token formData string true "Token"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /login/check [post]
func LoginCheck(ctx *gin.Context) {
	var token Token
	ctx.ShouldBind(&token)

	_, err := govalidator.ValidateStruct(token)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(401, gin.H{
			"code": -1,
			"data": "",
			"msg":  "检查授权失败",
		})
		return
	}

	claims, err := util.ParseToken(token.Token)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(401, gin.H{
			"code": -1,
			"data": "",
			"msg":  "检查授权失败",
		})
		return
	} else if time.Now().Unix() > claims.ExpiresAt {
		ctx.JSON(401, gin.H{
			"code": -1,
			"data": "",
			"msg":  "token已经过期",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code": 0,
		"data": map[string]string{"token": token.Token},
		"msg":  "检查授权成功",
	})
	return
}
