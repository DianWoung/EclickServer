package internal

import (
	"fmt"
)

type Room struct{
	roomID uint
	roomPlayers map[string]*User
}

func newRoom(roomID uint) *Room{
	fmt.Println("newRoom",roomID)
	room := Room{
		roomID:roomID,
		roomPlayers:map[string]*User{},
		}

	return  &room
}

