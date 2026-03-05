package engine

import (
	"sync"
	"github.com/gorilla/websocket"
)

// Player 玩家基础结构
type Player struct {
	ID       string
	Name     string
	Conn     *websocket.Conn
	IsReady  bool
	RoomID   string
	IsOnline bool // 标记在线状态
	Mu       sync.Mutex
}

// SendMessage 发送消息给玩家
func (p *Player) SendMessage(msg interface{}) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	if p.Conn != nil {
		p.Conn.WriteJSON(msg)
	}
}

// Table 桌子/房间结构
type Table struct {
	ID         string
	MaxPlayers int
	Players    []*Player
	GameLogic  interface{} // 动态接入不同游戏逻辑
	Status     int         // 0: 等待中, 1: 游戏中, 2: 结算中
	Mu         sync.RWMutex
}

// Broadcast 广播消息给桌上所有人
func (t *Table) Broadcast(msg interface{}) {
	t.Mu.RLock()
	defer t.Mu.RUnlock()
	for _, p := range t.Players {
		p.SendMessage(msg)
	}
}

// RoomManager 房间管理器
type RoomManager struct {
	Tables map[string]*Table
	Mu     sync.RWMutex
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		Tables: make(map[string]*Table),
	}
}

func (rm *RoomManager) CreateTable(id string, maxPlayers int) *Table {
	rm.Mu.Lock()
	defer rm.Mu.Unlock()
	t := &Table{
		ID:         id,
		MaxPlayers: maxPlayers,
		Players:    make([]*Player, 0),
	}
	rm.Tables[id] = t
	return t
}
