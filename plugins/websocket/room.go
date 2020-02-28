package websocket

import (
	"log"

	"github.com/google/uuid"
)

// Room struct
type Room struct {
	ID string
}

// NewRoom func
func NewRoom(pool *Pool) *Room {
	id, err := uuid.NewRandom()

	if err != nil {
		log.Println(err)
	}

	if len(pool.Rooms) != 0 {
		room := pool.Rooms[0]

		pool.RoomUnregister <- room

		return room
	}

	room := &Room{
		ID: id.String(),
	}

	pool.RoomRegister <- room

	return room
}
