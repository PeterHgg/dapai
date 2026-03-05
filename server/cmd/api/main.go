package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/user/dapai/internal/socket"
)

func main() {
	r := gin.Default()
	hub := socket.NewHub()
	go hub.Run()

	r.GET("/ws", func(c *gin.Context) {
		socket.WSHandler(hub, c)
	})

	log.Println("打牌服务端启动在 :8080")
	r.Run(":8080")
}
