package game

import (
	"gin-game/server/msg"
)

func Init() {
	msg.RegisterHandler(&msg.C2S_Login{}, handleLogin)
	msg.RegisterHandler(&msg.C2S_Register{}, handleRegister)
	msg.RegisterHandler(&msg.C2S_CreateRoom{}, handleCreateRoom)
}
