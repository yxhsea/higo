package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {
	goId := 123456
	merchantId := 20000332793
	integral := 10
	appKey := 20000332793
	appSecret := "508185bb88893454e8037e9079944a735b6defb1"
	timestamp := time.Now().Unix()
	fmt.Println("timestamp : ", timestamp)

	dataMap := map[string]interface{}{
		"app_key":     appKey,
		"app_secret":  appSecret,
		"merchant_id": merchantId,
		"go_id":       goId,
		"integral":    integral,
		"timestamp":   timestamp,
	}

	str, _ := json.Marshal(dataMap)
	fmt.Println(string(str))
	fmt.Println(getMapSort(dataMap))

	data := []byte(getMapSort(dataMap))
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	fmt.Println(md5str1)
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
