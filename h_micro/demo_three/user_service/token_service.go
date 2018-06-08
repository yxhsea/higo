package user_service

import (
	"github.com/dgrijalva/jwt-go"
	proto "higo/h_micro/demo_three/user_service/proto"
	"time"
)

type AuthAbel interface {
	Decode(tokenStr string) (*CustomClaims, error)
	Encode(user *proto.User) (string, error)
}

var privateKey = []byte("123456")

type CustomClaims struct {
	User               *proto.User
	jwt.StandardClaims //使用标准的payload
}

type TokenService struct {
}

//将Jwt字符串解密为CustomClaims对象
func (srv *TokenService) Decode(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})

	//解密转换类型并返回
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

//将User信息加密为Jwt字符串
func (srv *TokenService) Encode(user *proto.User) (string, error) {
	//三天后过期
	expireTime := time.Now().Add(time.Hour * 24 * 3).Unix()
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			Issuer:    "go.micro.srv.user", //签发者
			ExpiresAt: expireTime,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return jwtToken.SignedString(privateKey)
}
