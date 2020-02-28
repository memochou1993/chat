package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Client struct
type Client struct {
	ID   string
	Room *Room
	Conn *websocket.Conn
	Pool *Pool
}

// Message struct
type Message struct {
	RoomID   string `json:"roomId"`
	ClientID string `json:"clientId"`
	Type     int    `json:"type"`
	Body     string `json:"body"`
}

// NewClient func
func NewClient(pool *Pool, conn *websocket.Conn, room *Room, clientID string) *Client {
	return &Client{
		ID:   clientID,
		Room: room,
		Conn: conn,
		Pool: pool,
	}
}

// Read func
func (c *Client) Read() {
	defer func() {
		c.Pool.ClientUnregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		message := Message{
			RoomID:   c.Room.ID,
			ClientID: c.ID,
			Type:     messageType,
			Body:     string(p),
		}

		c.Pool.Broadcast <- message

		fmt.Printf("Message Received: %+v\n", message)
	}
}
