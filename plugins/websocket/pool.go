package websocket

import (
	"log"
)

// Pool struct
type Pool struct {
	Rooms            []*Room
	RoomRegister     chan *Room
	RoomUnregister   chan *Room
	Clients          map[*Client]bool
	ClientRegister   chan *Client
	ClientUnregister chan *Client
	Broadcast        chan Message
}

// NewPool func
func NewPool() *Pool {
	return &Pool{
		Rooms:            []*Room{},
		RoomRegister:     make(chan *Room),
		RoomUnregister:   make(chan *Room),
		Clients:          make(map[*Client]bool),
		ClientRegister:   make(chan *Client),
		ClientUnregister: make(chan *Client),
		Broadcast:        make(chan Message),
	}
}

// Start struct
func (pool *Pool) Start() {
	for {
		select {
		case room := <-pool.RoomRegister:
			pool.Rooms = append(pool.Rooms, room)
			break

		case <-pool.RoomUnregister:
			pool.Rooms = pool.Rooms[1:]
			break

		case client := <-pool.ClientRegister:
			pool.Clients[client] = true
			message := &Message{
				Type: 1,
				Body: "User has joined the conversation.",
			}
			notify(pool, client, message)
			break

		case client := <-pool.ClientUnregister:
			delete(pool.Clients, client)
			message := &Message{
				Type: 1,
				Body: "User has left the conversation.",
			}
			notify(pool, client, message)
			break

		case message := <-pool.Broadcast:
			broadcast(pool, message)
		}
	}
}

func notify(pool *Pool, c *Client, message *Message) {
	for client := range pool.Clients {
		if client.Room.ID == c.Room.ID && client.ID != c.ID {
			if err := client.Conn.WriteJSON(message); err != nil {
				log.Println(err)
			}
		}
	}
}

func broadcast(pool *Pool, message Message) {
	for client := range pool.Clients {
		if client.Room.ID == message.RoomID {
			if err := client.Conn.WriteJSON(message); err != nil {
				log.Println(err)
			}
		}
	}
}
