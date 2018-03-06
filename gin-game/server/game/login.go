package game

import (
	"gin-game/server/co"
	"gin-game/server/msg"
	"math/rand"

	"github.com/gorilla/websocket"
	"github.com/name5566/leaf/log"
)

func handleRegister(a *websocket.Conn, data []byte) {
	m := &msg.C2S_Register{}
	if !co.ConvertMsg(data, m) {
		return
	}

}

func handleLogin(a *websocket.Conn, data []byte) {
	m := &msg.C2S_Login{}
	if !co.ConvertMsg(data, m) {
		return
	}

	log.Debug("玩家登录 OpenID %v", m.OpenID)
	if m.OpenID == "123" {
		log.Debug("login failed OpenID %v", m.OpenID)
		return
	}

	if m.Version != 1 {
		log.Debug("login version failed Version %v", m.Version)
		return
	}

	rsps := &msg.S2C_Login{PlayerID: rand.Intn(99999) + 100000}
	msg.LinkPlayer(a, rsps.PlayerID)
	msg.WriteMsg(rsps.PlayerID, rsps)
}
