package controller

import (
	"log"
	"net/http"

	"github.com/memochou1993/chat/plugins/websocket"
)

// Handler func
func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		log.Println(err)
	}

	go websocket.Handle(conn)

	log.Println("Successfully Connected...")
}
