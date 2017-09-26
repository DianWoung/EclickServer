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
