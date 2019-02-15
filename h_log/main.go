package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

type AHandler struct {
	fwriter *os.File
}

func NewAHandler() *AHandler {
	fwriter, err := os.OpenFile("a.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	h := &AHandler{
		fwriter: fwriter,
	}
	return h
}

func (h *AHandler) Write(p []byte) (n int, err error) {
	// 将日志信息, 写入具体的收集器
	return h.fwriter.Write(p)
}

type BHandler struct {
	fwriter *os.File
}

func NewBHandler() *BHandler {
	filename := fmt.Sprintf("b-%s.log", time.Now().Format("2006-01-02"))
	fwriter, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	h := &BHandler{
		fwriter: fwriter,
	}
	return h
}

func (h *BHandler) Write(p []byte) (n int, err error) {
	// 将日志信息, 写入具体的收集器
	return h.fwriter.Write(p)
}

func main() {
	log.SetPrefix("[higo]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	log.SetOutput(io.MultiWriter(os.Stdout, NewAHandler(), NewBHandler()))
	// log.Println("111")

	r := gin.New()
	r.Use(Logger())
	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		GetLogger(c).logPrintf("3333")

		// it would print: "12345"
		log.Println(example)
	})

	err := r.Run(":8081")
	if err != nil {
		panic(err)
	}
}

type HLog struct {
	RequestId string
}

func (h *HLog) SetRequestId(requestId string) {
	h.RequestId = requestId
}

func (h *HLog) logPrintf(params string) {
	logger := log.New(os.Stdout, "h_log", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	err := logger.Output(2, fmt.Sprintf("[requestId: %s], [params: %s]", h.RequestId, params))
	if err != nil {
		panic(err)
	}
}

func SetLogger(c *gin.Context) {
	hLog := &HLog{}
	hLog.SetRequestId(fmt.Sprintf("1111-2222-3333-%d", rand.Intn(100)))
	hLogBytes, err := json.Marshal(hLog)
	if err != nil {
		panic(err)
	}
	c.Set("hlog", string(hLogBytes))
}

func GetLogger(c *gin.Context) *HLog {
	hLogBytes, ok := c.Get("hlog")
	if !ok {
		panic(ok)
	}

	hLog := &HLog{}
	err := json.Unmarshal([]byte(fmt.Sprintf("%s", hLogBytes)), &hLog)
	if err != nil {
		panic(err)
	}

	return hLog
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 初始化 log 对象
		SetLogger(c)

		GetLogger(c).logPrintf("1111")

		t := time.Now()
		c.Set("example", "12345")

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		GetLogger(c).logPrintf("2222")

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
