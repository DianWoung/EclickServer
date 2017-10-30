package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"server/msg"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), func(args []interface{}) {
		// user
		//监测用户是否登录
		a := args[1].(gate.Agent)
		user := users[a.UserData().(*AgentInfo).userID]
		if user == nil {
			a.WriteMsg(&msg.S2C_Close{Err: msg.S2C_Auth_AccIDInvalid})
			log.Debug("This agent is not registered")
			return
		}

		// agent to user
		args[1] = user
		h.(func([]interface{}))(args)
	})
}

func init() {
	handleMsg(&msg.C2C_Msg{}, handleC2CMsg)
	handleMsg(&msg.C2S_CreateRoom{},handleCreateRoom)
	handleMsg(&msg.C2S_JoinRoom{},handleJoinRoom)
	handleMsg(&msg.C2S_LeftRoom{},handleLeftRoom)
	handleMsg(&msg.C2S_BroadcastRoom{},handleBroadcastRoom)
	handleMsg(&msg.C2S_GetRooms{},handleGetRooms)
	handleMsg(&msg.C2S_RemoveRoom{},handleRemoveRoom)
}

func handleC2CMsg(args []interface{}) {
	m := args[0].(*msg.C2C_Msg)
	a := args[1].(gate.Agent)
	//目标主机
	client := accIDUsers[m.AccID]
	if client == nil {
		a.WriteMsg(&msg.S2C_Close{Err: msg.S2C_Close_InnerError})
		return
	}
	client.WriteMsg(&msg.S2C_Msg{
		m.X,
		m.Y,
	})
}

func handleCreateRoom(args []interface{})  {
	a := args[1].(gate.Agent)
	//
	room := CreateRoom(a.UserData().(*AgentInfo).accID)
	if room == nil {
		a.WriteMsg(&msg.S2C_Room{
			msg.S2C_Room_UnknownErr,
		})
		return
	}
	a.WriteMsg(&msg.S2C_NewRoomID{
		room.roomID,
	})
}

func handleJoinRoom(args []interface{})  {
	m := args[0].(*msg.C2S_JoinRoom)
	a := args[1].(gate.Agent)
	ok := JoinRoom(m.ID,a.UserData().(*AgentInfo).accID)
	a.WriteMsg(&msg.S2C_Room{ok})

}

func handleLeftRoom(args []interface{})  {
	m := args[0].(*msg.C2S_LeftRoom)
	a := args[1].(gate.Agent)
	errCode := LeftRoom(m.ID,a.UserData().(*AgentInfo).accID)
	a.WriteMsg(&msg.S2C_Room{errCode})

}

func handleBroadcastRoom(args []interface{})  {
	m := args[0].(*msg.C2S_BroadcastRoom)
	a := args[1].(gate.Agent)
	ok :=BroadcastRoom(m.ID,m.Msg,a.UserData().(*AgentInfo).accID)
	if ok!=msg.S2C_Room_OK {
		a.WriteMsg(&msg.S2C_Room{ok})
	}
}

func handleGetRooms(args []interface{})  {
	a := args[1].(gate.Agent)
	rooms := GetAllRooms()
	roomList := []msg.RoomInfo{}
	for _,v := range rooms {
	roomList = append(roomList,msg.RoomInfo{v.roomID})
	}
	a.WriteMsg(&msg.S2C_AllRooms{
		RoomList:roomList,
	})
}

func handleRemoveRoom(args []interface{})  {
	m := args[0].(*msg.C2S_BroadcastRoom)
	a := args[1].(gate.Agent)
	ok := RemoveRoom(m.ID)
	if ok!=msg.S2C_Room_OK {
		a.WriteMsg(&msg.S2C_Room{ok})
	}
}