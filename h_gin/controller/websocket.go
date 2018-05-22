package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

//WebSocket
func WebSocket(ctx *gin.Context) {
	if ctx.IsWebsocket() {
		fmt.Println("WebSocket connection...")
	}

	Upgrader := websocket.Upgrader{
		ReadBufferSize:  int(1024),
		WriteBufferSize: int(1024),
		CheckOrigin:     func(r *http.Request) bool { return true },
		Subprotocols:    []string{"binary"},
	}

	ws, err := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Printf("WebSocket upgrade failed : %v \n", err)
	}

	ws.Close()
}
