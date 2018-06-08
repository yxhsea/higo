package handle

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	proto "higo/h_rpc/user_service/proto"
	"time"
)

type AuthAbel interface {
	Decode(tokenStr string) (*CustomClaims, error)
	Encode(user *proto.User) (string, error)
}

var privateKey = []byte("1234abc")

type CustomClaims struct {
	Ip                 string `json:"ip"`
	Password           string `json:"password"`
	jwt.StandardClaims        //使用标准的payload
}

type TokenService struct {
}

//将Jwt字符串解密为CustomClaims对象
func (srv *TokenService) Decode(tokenStr string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})

	if tokenClaims != nil {
		//解密转换类型并返回
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		} else {
			return nil, err
		}
	}
	return nil, err
}

//将User信息加密为Jwt字符串
func (srv *TokenService) Encode(user *proto.User) (string, error) {
	//三天后过期
	expireTime := time.Now().Add(time.Hour * 24 * 3).Unix()
	claims := CustomClaims{
		encodeMD5(user.Ip),
		encodeMD5(user.Password),
		jwt.StandardClaims{
			Issuer:    "go.micro.srv.user", //签发者
			ExpiresAt: expireTime,
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(privateKey)
}

func encodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
