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
}
