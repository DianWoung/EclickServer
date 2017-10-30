package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&S2C_Close{})
	Processor.Register(&C2S_Auth{})
	Processor.Register(&S2C_Auth{})
	Processor.Register(&C2C_Msg{})
	Processor.Register(&S2C_Msg{})
	Processor.Register(&C2S_JoinRoom{})
	Processor.Register(&C2S_LeftRoom{})
	Processor.Register(&C2S_CreateRoom{})
	Processor.Register(&C2S_BroadcastRoom{})
	Processor.Register(&C2S_GetRooms{})
	Processor.Register(&S2C_NewRoomID{})
	Processor.Register(&S2C_AllRooms{})
	Processor.Register(&S2C_Room{})
	Processor.Register(&S2C_MsgFromRoom{})
	Processor.Register(&C2S_RemoveRoom{})
}

// Close
const (
	S2C_Close_LoginRepeated = 1
	S2C_Close_InnerError    = 2
)

type S2C_Close struct {
	Err int
}

// Auth
type C2S_Auth struct {
	AccID string
}

const (
	S2C_Auth_OK           = 0
	S2C_Auth_AccIDInvalid = 1
)

type S2C_Auth struct {
	Err int
}

type C2C_Msg struct {
	AccID string
	X     int
	Y     int
}

type S2C_Msg struct {
	X int
	Y int
}
// room modules
const (
	S2C_Room_OK = 0
	S2C_Room_NotMember = 1
	S2C_Room_IdInvalid = 2
	S2C_Room_SendFailed = 3
	S2C_Room_UnknownErr = 4
	S2C_Room_AccIDInvalid  = 5
)
type C2S_CreateRoom struct {

}

type C2S_JoinRoom struct {
	ID uint `json:"id"`
}

type C2S_LeftRoom struct {
	ID uint `json:"id"`
}

type C2S_BroadcastRoom struct {
	ID uint `json:"id"`
	Msg string `json:"msg"`
}

type C2S_GetRooms struct {

}

type S2C_AllRooms struct {
	RoomList []RoomInfo `json:"roomlist"`
}
type RoomInfo struct {
	ID uint `json:"id"`
}

type S2C_NewRoomID struct {
	ID uint `json:"id"`
}

type S2C_Room struct {
	Err uint `json:"err"`
}

type S2C_MsgFromRoom struct {
	AccID string `json:"accid"`
	Msg string `json:"msg"`
}

type C2S_RemoveRoom struct {
	ID uint `json:"id"`
}