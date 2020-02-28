package controller

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/memochou1993/chat/helper"
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

	clientID := base64.StdEncoding.EncodeToString([]byte(helper.GetHost(r)))

	room := websocket.NewRoom(pool, clientID)

	client := websocket.NewClient(pool, conn, room, clientID)

	pool.ClientRegister <- client

	client.Read()
}
