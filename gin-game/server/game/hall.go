package game

import (
	"gin-game/server/co"
	"gin-game/server/def"
	"gin-game/server/msg"
	"math/rand"

	"github.com/name5566/leaf/log"

	"github.com/gorilla/websocket"
)

func handleCreateRoom(a *websocket.Conn, data []byte) {
	m := &msg.C2S_CreateRoom{}
	if !co.ConvertMsg(data, m) {
		return
	}

	playerID := msg.GetPlayer(a)
	if playerID == def.InvalidPlayerID {
		log.Debug("not login")
		return
	}

	log.Debug("handleCreateRoom(%d)", playerID)

	msg.WriteMsg(playerID, &msg.S2C_CreateRoom{RoomKey: rand.Intn(100) + 100000})
}
