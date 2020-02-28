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
		Rooms:            []*Room{},        // New
		RoomRegister:     make(chan *Room), // New
		RoomUnregister:   make(chan *Room), // New
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
			pool.Rooms = pool.Rooms[1:len(pool.Rooms)]
			fmt.Println("Number of rooms: ", len(pool.Rooms))
			break

		case client := <-pool.ClientRegister:
			pool.Clients[client] = true
			fmt.Println("Number of clients: ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Connected..."})
			}
			break

		case client := <-pool.ClientUnregister:
			delete(pool.Clients, client)
			fmt.Println("Number of clients: ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
			break

		case message := <-pool.Broadcast:
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}
