package websocket

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/memochou1993/chat/helper"
)

// Client struct
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

// Message struct
type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

// NewClient func
func NewClient(r *http.Request, conn *websocket.Conn, pool *Pool) *Client {
	return &Client{
		ID:   base64.StdEncoding.EncodeToString([]byte(helper.GetHost(r))),
		Conn: conn,
		Pool: pool,
	}
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		message := Message{Type: messageType, Body: string(p)}

		c.Pool.Broadcast <- message

		fmt.Printf("Message Received: %+v\n", message)
	}
}
