package internal

import (
	"github.com/name5566/leaf/log"
	"server/msg"
)
var(
	rooms = make(map[uint]*Room)
	lastRoomID = uint(0)
)

func CreateRooms(roomID uint) *Room{
	room := newRoom(roomID)
	rooms[roomID] = room
	lastRoomID ++
	return  room
}


func GetRoomFromID(roomID uint) *Room{
	room, ok := rooms[roomID]
	if(!ok){
		return nil
	}
	return room
}

func JoinRoom(roomID uint,accID string) uint {
	room,ok := rooms[roomID]
	if(!ok){
		return msg.S2C_Room_IdInvalid
	}
	isIn := room.roomPlayers[accID]
	if isIn !=nil {
		log.Debug("is is nil")
		return msg.S2C_Room_AccIDInvalid
	}
	user := accIDUsers[accID]
	room.roomPlayers[accID] = user
	return msg.S2C_Room_OK
}

func GetAllRooms() map[uint]*Room {
	return rooms
}

func RemoveRoom(roomID uint) uint  {
	_,ok := rooms[roomID]
	if (!ok) {
		log.Debug("unknown roomid:%s",roomID)
		return msg.S2C_Room_IdInvalid
	}
	delete(rooms,roomID)
	return msg.S2C_Room_OK
}


func BroadcastRoom(roomID uint,message string,accID string) uint {
	room,ok := rooms[roomID]
	if (!ok) {
		return msg.S2C_Room_IdInvalid
	}
	isIn := room.roomPlayers[accID]
	if isIn == nil {
		return msg.S2C_Room_NotMember
	}
	for _,a := range room.roomPlayers {
		a.WriteMsg(&msg.S2C_MsgFromRoom{
			AccID:accID,
			Msg:message,
		})
	}
	return msg.S2C_Room_OK
}

//room模块

func CreateRoom(accID string) *Room {
	newRoom := CreateRooms(lastRoomID)
	ok :=JoinRoom(newRoom.roomID,accID)
	if ok==msg.S2C_Room_OK {
		return newRoom
	}
	return nil
}

func LeftRoom(roomID uint,accID string) uint  {
	room,ok := rooms[roomID]
	if (!ok) {
		log.Debug("unknown roomid",roomID)
		return msg.S2C_Room_IdInvalid
	}
	isIn := room.roomPlayers[accID]
	if isIn == nil {
		return msg.S2C_Room_NotMember
	}
	delete(room.roomPlayers,accID)
	if len(room.roomPlayers) == 0 {
		RemoveRoom(roomID)
	}
	return msg.S2C_Room_OK
}