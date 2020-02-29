package websocket

import (
	"fmt"
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
			fmt.Println("Number of rooms: ", len(pool.Rooms))
			break

		case <-pool.RoomUnregister:
			pool.Rooms = pool.Rooms[1:]
			fmt.Println("Number of rooms: ", len(pool.Rooms))
			break

		case client := <-pool.ClientRegister:
			message := &Message{
				ClientID: client.ID,
				Type:     1,
				Body:     "User has joined the conversation.",
			}
			pool.Clients[client] = true
			notify(pool, client, message)
			break

		case client := <-pool.ClientUnregister:
			message := &Message{
				ClientID: client.ID,
				Type:     1,
				Body:     "User has left the conversation.",
			}
			notify(pool, client, message)
			delete(pool.Clients, client)
			break

		case message := <-pool.Broadcast:
			broadcast(pool, message)
		}
	}
}

func notify(pool *Pool, c *Client, message *Message) {
	self := 0

	for client := range pool.Clients {
		if client.ID == c.ID {
			self++
		}

		if self > 1 {
			return
		}
	}

	for client := range pool.Clients {
		if client.Room.ID != c.Room.ID {
			continue
		}

		if err := client.Conn.WriteJSON(message); err != nil {
			log.Println(err)
		}
	}
}

func broadcast(pool *Pool, message Message) {
	for client := range pool.Clients {
		if client.Room.ID != message.RoomID {
			continue
		}

		if err := client.Conn.WriteJSON(message); err != nil {
			log.Println(err)
		}
	}
}
