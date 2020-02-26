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

	client := websocket.NewClient(r, conn, pool)

	pool.Register <- client

	client.Read()
}
