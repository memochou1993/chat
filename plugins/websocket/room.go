package websocket

import (
	"github.com/memochou1993/chat/helper"
)

// Room struct
type Room struct {
	ID string
}

// NewRoom func
func NewRoom(pool *Pool, clientID string) *Room {
	if room := find(pool, clientID); room != nil {
		return room
	}

	if room := assign(pool, clientID); room != nil {
		return room
	}

	room := create(pool, clientID)

	return room
}

func find(pool *Pool, clientID string) *Room {
	for client := range pool.Clients {
		if client.ID == clientID {
			return client.Room
		}
	}

	return nil
}

func assign(pool *Pool, clientID string) *Room {
	if len(pool.Rooms) != 0 {
		room := pool.Rooms[0]

		pool.RoomUnregister <- room

		return room
	}

	return nil
}

func create(pool *Pool, clientID string) *Room {
	room := &Room{
		ID: getRoomID(),
	}

	pool.RoomRegister <- room

	return room
}

func getRoomID() string {
	return helper.GetUUID()
}
