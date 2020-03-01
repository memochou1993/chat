package controller

import (
	"log"
	"net/http"

	"github.com/memochou1993/chat/plugins/websocket"
)

var (
	pool = websocket.NewPool()
)

func init() {
	go pool.Start()
}

// Handler func
func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		log.Println(err)
	}

	clientID := websocket.GetClientID(r)

	room := websocket.NewRoom(pool, clientID)

	client := websocket.NewClient(pool, conn, room, clientID)

	pool.ClientRegister <- client

	client.ReadMessage()
}
