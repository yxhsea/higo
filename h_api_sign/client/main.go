package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
	"time"
)

const AccessKey = "Ztp4qYYOdLxWbQ"                   //开发者标识，确保唯一
const SecretKey = "BnV0RvocoZZsvRrjS5L22XPbCWOtqFts" //用于接口加密，确保不易被枚举，生成算法不易被猜测

func main() {
	/** timestamp+nonce 时间戳加随机字符串解决方案
	nonce指唯一的随机字符串，用来标识每个被签名的请求。通过为每个请求提供一个唯一的标识符，服务器能够防止请求被多次使用（记录所有用过的nonce以阻止它们被二次使用）。
	然而，对服务器来说永久存储所有接收到的nonce的代价是非常大的。可以使用timestamp来优化nonce的存储。
	假设允许客户端和服务端最多能存在15分钟的时间差，同时追踪记录在服务端的nonce集合。
	当有新的请求进入时，首先检查携带的timestamp是否在2分钟内，如超出时间范围，则拒绝，然后查询携带的nonce，如存在已有集合，则拒绝。
	否则，记录该nonce，并删除集合内时间戳大于2分钟的nonce（可以使用redis的expire，新增nonce的同时设置它的超时失效时间为15分钟）。
	*/

	//timestamp 时间戳、 nonce_str 随机字符串

	userId := "123456"                             //用户ID
	amount := "10"                                 //金额
	nonceStr := "sJxWXV66iy6pD2K0jrhS1S9yLGGQrhx4" //随机字符串
	nowTime := time.Now().Unix()                   //时间戳
	dataMap := map[string]interface{}{
		"access_key": AccessKey,
		"user_id":    userId,
		"amount":     amount,
		"timestamp":  nowTime,
		"nonce_str":  nonceStr,
	}
	dataMap["sign"] = getSign(dataMap)

	fmt.Printf("user_id:%v\n", dataMap["user_id"])
	fmt.Printf("amount:%v\n", dataMap["amount"])
	fmt.Printf("access_key:%v\n", dataMap["access_key"])
	fmt.Printf("timestamp:%v\n", dataMap["timestamp"])
	fmt.Printf("nonce_str:%v\n", dataMap["nonce_str"])
	fmt.Printf("sign:%v\n", dataMap["sign"])
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
