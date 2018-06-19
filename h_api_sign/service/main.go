package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

const AccessKey = "Ztp4qYYOdLxWbQ"                   //开发者标识，确保唯一
const SecretKey = "BnV0RvocoZZsvRrjS5L22XPbCWOtqFts" //用于接口加密，确保不易被枚举，生成算法不易被猜测

var RedisConn *redis.Client

func main() {
	//redis init.
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.0.25:6379",
		Password: "",
		DB:       2,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Redis init fail, Error : %v ", err.Error())
	}
	RedisConn = client

	//http init.
	router := gin.Default()
	router.POST("/test", handler)
	router.Run(":8887")
}

func handler(ctx *gin.Context) {
	accessKey := ctx.PostForm("access_key")
	timestamp := ctx.PostForm("timestamp")
	nonceStr := ctx.PostForm("nonce_str")
	sign := ctx.PostForm("sign")

	userId := ctx.PostForm("user_id")
	amount := ctx.PostForm("amount")

	//判断签名有没有过期，过期时间为20s
	times, _ := com.StrTo(timestamp).Int64()
	nowTime := time.Now().Unix()
	if nowTime-times > 120 {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "签名过期"})
		return
	}

	//判断签名是否有误
	dataMap := map[string]interface{}{
		"access_key": accessKey,
		"user_id":    userId,
		"amount":     amount,
		"timestamp":  timestamp,
		"nonce_str":  nonceStr,
	}
	if sign != getSign(dataMap) {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "签名错误"})
		return
	}

	//判断是否存在nonce_str随机字符串
	flag, _ := RedisConn.Exists(nonceStr).Result()
	if flag > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "签名失效"})
		return
	}

	//存储nonce_str随机字符串
	valInt64, _ := RedisConn.Set(nonceStr, timestamp, time.Second*120).Result()
	fmt.Println(valInt64)

	ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
	return
}

func getMapSort(sortMap map[string]interface{}) string {
	keys := make([]string, len(sortMap))
	i := 0
	for k := range sortMap {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, k := range keys {
		if k != "sign" && !strings.HasPrefix(k, "reserved") {
			buf.WriteString(k)
			buf.WriteString("=")
			buf.WriteString(fmt.Sprint(sortMap[k]))
			buf.WriteString("&")
		}
	}
	bufStr := buf.String()
	return strings.TrimRight(bufStr, "&")
}

func getSign(dataMap map[string]interface{}) string {
	dataMap["secret_key"] = SecretKey
	data := []byte(getMapSort(dataMap))
	delete(dataMap, "secret_key")
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return strings.ToUpper(md5str)
}
