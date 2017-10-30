package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/go"
	"github.com/satori/go.uuid"
)

var (
	accIDUsers = make(map[string]*User)
	users      = make(map[string]*User)
)

const (
	userLogin = iota
	userLogout
	userGame
)

type User struct {
	gate.Agent
	*g.LinearContext
	state int
	data  *UserData
}

func (user *User) login(accID string) {
	userData := new(UserData)
	skeleton.Go(func() {
		userData.UserID = uuid.NewV4().String()
	}, func() {
		// network closed
		if user.state == userLogout {
			user.logout(accID)
			return
		}

		// db error
		user.state = userGame
		if userData == nil {
			return
		}

		// ok
		user.data = userData
		users[userData.UserID] = user
		user.UserData().(*AgentInfo).userID = userData.UserID
		user.onLogin()
	})
}

func (user *User) logout(accID string) {
	if user.data != nil {
		user.onLogout()
		delete(users, user.data.UserID)
	}

	// save
	//data := util.DeepClone(user.data)
	user.Go(func() {

	}, func() {
		delete(accIDUsers, accID)
	})
}

func (user *User) isOffline() bool {
	return user.state == userLogout
}

func (user *User) onLogin() {

}

func (user *User) onLogout() {

}

