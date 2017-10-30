package gate

import (
	"server/game"
	"server/login"
	"server/msg"
)

func init() {
	// login
	msg.Processor.SetRouter(&msg.C2S_Auth{}, login.ChanRPC)

	// game
	msg.Processor.SetRouter(&msg.C2C_Msg{}, game.ChanRPC)

	//room
	msg.Processor.SetRouter(&msg.C2S_CreateRoom{},game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_JoinRoom{},game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_LeftRoom{},game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_BroadcastRoom{},game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_GetRooms{},game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_RemoveRoom{},game.ChanRPC)
}
