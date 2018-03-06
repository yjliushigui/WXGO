package game

import "gin-game/server/msg"

// 在线状态
const (
	NSPlayerOffline = 0
	NSPlayerOnline  = 1
)

type tPlayerCache struct {
	// 网络状态 0-离线 1-服务器主动关闭连接 2-在线
	netState int
	// 所在房间号 0-没有在房间中
	roomKey int
	// 经度
	Longitude float32
	// 纬度信息
	Latitude float32
}

var (
	playercache = make(map[int]*tPlayerCache)
)

func getPlayerCache(playerID int) *tPlayerCache {
	return playercache[playerID]
}

func createPlayerCache(playerID int) *tPlayerCache {
	var ret *tPlayerCache
	if c, exist := playercache[playerID]; exist {
		// 检查玩家是否在房间中，如果在则检查是否是断线重连成功
		rcache := getRoomCache(c.roomKey)
		if rcache != nil && rcache.JoinRoom(playerID) {
			c.netState = NSPlayerOnline
			rcache.Broadcast(&msg.S2C_Online{PlayerID: playerID})
		}
		ret = c
	} else {
		ret = &tPlayerCache{netState: NSPlayerOnline}
		playercache[playerID] = ret
	}

	return ret
}

func (c *tPlayerCache) SetRoomKey(roomKey int) {
	c.roomKey = roomKey
}

func (c *tPlayerCache) Offline() {
	c.netState = NSPlayerOffline
}

func (c *tPlayerCache) Online() {
	c.netState = NSPlayerOnline
}

func (c *tPlayerCache) State() int {
	return c.netState
}
