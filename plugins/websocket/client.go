package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
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
func NewClient(r *http.Request, room *Room, conn *websocket.Conn, pool *Pool) *Client {
	id, err := uuid.NewRandom()

	if err != nil {
		log.Println(err)
	}

	return &Client{
		// ID:   base64.StdEncoding.EncodeToString([]byte(helper.GetHost(r))),
		ID:   id.String(),
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
