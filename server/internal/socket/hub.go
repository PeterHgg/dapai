package socket

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 微信浏览器开发环境可设为 true
	},
}

// Hub 处理 WebSocket 连接分发
type Hub struct {
	// 注册、注销、广播通道
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
}

type Client struct {
	ID   string
	Conn *websocket.Conn
	Hub  *Hub
	Send chan []byte
}

func NewHub() *Hub {
	return &Hub{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			log.Printf("新连接注册: %s", client.ID)
		case client := <-h.Unregister:
			log.Printf("连接注销: %s", client.ID)
		case message := <-h.Broadcast:
			log.Printf("全局广播消息: %s", string(message))
		}
	}
}

// WSHandler 处理 HTTP 到 WebSocket 的升级
func WSHandler(hub *Hub, c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("升级连接失败:", err)
		return
	}

	client := &Client{
		ID:   c.Query("uid"), // 简易 UID
		Conn: conn,
		Hub:  hub,
		Send: make(chan []byte, 256),
	}
	hub.Register <- client

	// 这里通常启动两个协程处理读写
	go client.ReadPump()
}

func (cl *Client) ReadPump() {
	defer func() {
		cl.Hub.Unregister <- cl
		cl.Conn.Close()
	}()
	for {
		_, _, err := cl.Conn.ReadMessage()
		if err != nil {
			break
		}
		// 处理收到消息的逻辑
	}
}
