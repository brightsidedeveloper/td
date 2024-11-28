package logic

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type GameWebSocket struct {
	Logic *GameLogic
	mu    sync.Mutex
	Clients map[*websocket.Conn]bool
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (gws *GameWebSocket) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Error upgrading websocket: %v\n", err)
		return
	}
	gws.mu.Lock()
	gws.Clients[conn] = true
	gws.mu.Unlock()

	defer func() {
		gws.mu.Lock()
		delete(gws.Clients, conn)
		gws.mu.Unlock()
		conn.Close()
	}()

	for {
		gameState := gws.Logic.State()
		if err := conn.WriteJSON(gameState); err != nil {
			fmt.Printf("Error writing JSON: %v\n", err)
			break
		}
	}
}

func NewGameWebSocket(logic *GameLogic) *GameWebSocket {
	return &GameWebSocket{
		Logic: logic,
		Clients: make(map[*websocket.Conn]bool),
	}
}