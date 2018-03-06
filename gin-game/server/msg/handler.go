package msg

import (
	"encoding/json"
	"gin-game/server/co"
	"gin-game/server/def"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/name5566/leaf/log"
)

var (
	handlers   map[string]func(a *websocket.Conn, m []byte)
	wsupgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	connmgr     map[*websocket.Conn]int
	player2conn map[int]*websocket.Conn
)

func init() {
	handlers = make(map[string]func(a *websocket.Conn, m []byte))
	connmgr = make(map[*websocket.Conn]int)
	player2conn = make(map[int]*websocket.Conn)
}

func RegisterHandler(m interface{}, h func(a *websocket.Conn, m []byte)) {
	handlers[co.TypeName(m)] = h
}

func dispatch(id string, a *websocket.Conn, m []byte) {
	log.Debug("handle1: %v", id)
	h, exist := handlers[id]
	if !exist {
		log.Error("no dispatch, id=%s, handlers=%v", id, handlers)
		return
	}

	h(a, m)
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	wsupgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("Failed to set websocket upgrade %v", err)
		return
	}

	if _, exist := connmgr[conn]; !exist {
		connmgr[conn] = def.InvalidPlayerID
	}

	for {
		_, m, err := conn.ReadMessage()
		if err != nil {
			break
		}

		defer func() {
			if err := recover(); err != nil {
				log.Error("dispatch recover %v", err)
			}
		}()

		// log.Debug("recv", "msg", m)

		// 解析协议
		var raw map[string]json.RawMessage
		err = json.Unmarshal(m, &raw)
		if err != nil {
			log.Error("protocol umarshal failed %v", err)
			return
		}

		for name, data := range raw {
			dispatch(name, conn, data)
		}
	}
}

func LinkPlayer(conn *websocket.Conn, playerID int) {
	if _, exist := connmgr[conn]; !exist {
		log.Debug("player(%d) agent无效", playerID)
		return
	}

	connmgr[conn] = playerID
	player2conn[playerID] = conn
}

func GetPlayer(conn *websocket.Conn) int {
	return connmgr[conn]
}

func GetAgent(playerID int) *websocket.Conn {
	return player2conn[playerID]
}

func UnlinkAgent(playerID int) {
	a := GetAgent(playerID)
	if a == nil {
		return
	}
	delete(connmgr, a)
	delete(player2conn, playerID)
}

func WriteMsg(playerID int, m interface{}) {
	a := GetAgent(playerID)
	if a == nil {
		return
	}

	a.WriteJSON(map[string]interface{}{co.TypeName(m): m})
	log.Debug("send2player %v", m)
}
