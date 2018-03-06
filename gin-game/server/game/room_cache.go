package game

import (
	"gin-game/server/def"
	"gin-game/server/msg"
	"time"

	"github.com/name5566/leaf/log"
)

// 投票解散房间
const (
	// VoteQuitRoomNone 没有发起投票
	VoteQuitRoomNone = 0
	// VoteQuitRoomNo 不同意解散
	VoteQuitRoomNo = 1
	// VoteQuitRoomOK 同意解散
	VoteQuitRoomOK = 2

	voteDisbanTimeout = 55
)

type tRoomCache struct {
	// 投票退房 0-没有发起投票, 1-不同意退房, 2-退房
	voteQuitRoom map[int]int
	// 发起投票时间
	voteStartTime time.Time
	// 投票申请人ID
	applyPlayerID int

	// 游戏实例
	gameInst def.IGame
	info     *msg.C2S_CreateRoom
	// 房间里的玩家
	players []int
}

var (
	roomcache = make(map[int]*tRoomCache)
)

const maxRoomPlayerCount = 3

func (c tRoomCache) IsFull() bool {
	return c.PlayerCount() >= maxRoomPlayerCount
}

func (c tRoomCache) GetOwner() int {
	if len(c.players) == 0 {
		return 0
	}

	return c.players[0]
}

func (c tRoomCache) Broadcast(m interface{}) {
	for _, playerIDn := range c.players {
		a := msg.GetAgent(playerIDn)
		if a == nil {
			log.Debug("player(%d) no agent", playerIDn)
			continue
		}

		a.WriteJSON(m)
	}
}

func (c tRoomCache) WriteOwnerMsg(m interface{}) {
	a := msg.GetAgent(c.GetOwner())
	if a == nil {
		log.Debug("player(%d) no agent", c.GetOwner())
		return
	}

	a.WriteJSON(m)
}

func (c tRoomCache) Players() []int {
	arr := make([]int, 0, 4)
	for _, playerIDn := range c.players {
		if playerIDn != def.InvalidPlayerID {
			arr = append(arr, playerIDn)
		}
	}
	return arr
}

func (c tRoomCache) PlayerCount() int {
	count := 0
	for _, playerIDn := range c.players {
		if playerIDn != def.InvalidPlayerID {
			count++
		}
	}

	return count
}

func (c tRoomCache) InRoom(playerID int) bool {
	for _, playerIDn := range c.players {
		if playerID == playerIDn && playerID != def.InvalidPlayerID {
			return true
		}
	}

	return false
}

func destroyRoom() {
	for roomKey, cache := range roomcache {
		if cache.gameInst != nil && cache.gameInst.State() == def.GamePlaying {
			cache.gameInst.Stop()
		}
		log.Release("stop room(%d)", roomKey)
	}
}

func (c *tRoomCache) JoinRoom(playerID int) bool {
	if c.PlayerCount() >= maxRoomPlayerCount && !c.InRoom(playerID) {
		// 房间玩家已满
		log.Debug("player(%d) join room full!", playerID)
		return false
	}

	if c.InRoom(playerID) {
		log.Release("player(%d) reenter in room", playerID)
	} else {
		var joined bool
		for idx, playerIDn := range c.players {
			if playerIDn == def.InvalidPlayerID {
				c.players[idx] = playerID
				joined = true
				break
			}
		}

		if !joined {
			c.players = append(c.players, playerID)
		}
	}

	return true
}

func (c *tRoomCache) LeaveRoom(playerID int) {
	for idx, playerIDn := range c.players {
		if playerID == playerIDn {
			// 将退出玩家位置留空
			c.players[idx] = 0
			break
		}
	}
}

func (c *tRoomCache) VoteDisbandReset() {
	c.voteStartTime = time.Time{}
	c.voteQuitRoom = nil
}

func (c tRoomCache) IsGaming() bool {
	return c.gameInst != nil && c.gameInst.State() == def.GamePlaying
}

func (c *tRoomCache) ClearRoomKey() {
	roomKey := 0
	c.ForeachPlayer(func(playerIDn int) {
		playerCacheN := getPlayerCache(playerIDn)
		roomKey = playerCacheN.roomKey
		playerCacheN.roomKey = def.InvalidRoomKey
	})
	delete(roomcache, roomKey)
}

func (c *tRoomCache) VoteDisbandTimeout() bool {
	diff := time.Since(c.voteStartTime).Seconds()
	log.Debug("VoteDisbandTimeout diff %f", diff)
	return diff > voteDisbanTimeout
}

func (c *tRoomCache) VoteDisbandRoom(playerID int) bool {
	if time.Since(c.voteStartTime).Seconds() > voteDisbanTimeout {
		c.voteStartTime = time.Now()
		c.voteQuitRoom = make(map[int]int)
		c.voteQuitRoom[playerID] = VoteQuitRoomOK
		// 记录申请投票人
		c.applyPlayerID = playerID
		log.Debug("记录首个申请投票人(%d)", playerID)
		return false
	}

	c.voteQuitRoom[playerID] = VoteQuitRoomOK
	disbandCount := 0
	for _, val := range c.voteQuitRoom {
		if val == VoteQuitRoomOK {
			disbandCount++
		}
	}

	if disbandCount >= c.PlayerCount() {
		// 所有人都同意解散房间
		return true
	}

	return false
}

func (c tRoomCache) ForeachVoted(fn func(playerID int)) {
	for playerIDn := range c.voteQuitRoom {
		fn(playerIDn)
	}
}

// func (c tRoomCache) VoteState() *msg.S2C_ContinueVote {
// 	rsps := &msg.S2C_ContinueVote{
// 		ApplyPlayerID: c.applyPlayerID,
// 		PlayerID:      make([]int, 0, len(c.voteQuitRoom)),
// 		Disband:       make([]int, 0, len(c.voteQuitRoom)),
// 	}

// 	for playerID, disband := range c.voteQuitRoom {
// 		rsps.PlayerID = append(rsps.PlayerID, playerID)
// 		rsps.Disband = append(rsps.Disband, disband)
// 	}

// 	return rsps
// }

func (c tRoomCache) IsVoting() bool {
	log.Debug("IsVoting: %v %v", c.voteQuitRoom, c.VoteDisbandTimeout())
	return c.gameInst != nil && c.gameInst.State() == def.GamePlaying && len(c.voteQuitRoom) > 0 && !c.VoteDisbandTimeout()
}

func (c tRoomCache) IsVoteTimeout() bool {
	disbandCount := 0
	for _, val := range c.voteQuitRoom {
		if val == VoteQuitRoomOK {
			disbandCount++
		}
	}

	log.Debug("disbandCount %d, room player %d", disbandCount, c.PlayerCount())
	return disbandCount != 0 && disbandCount < c.PlayerCount()
}

func (c tRoomCache) ForeachPlayer(fn func(playerID int)) {
	for _, playerIDn := range c.players {
		if playerIDn == def.InvalidPlayerID {
			continue
		}

		fn(playerIDn)
	}
}

func makeRoomCache(playerID, roomKey int, m *msg.C2S_CreateRoom) {
	players := make([]int, 0, maxRoomPlayerCount)
	players = append(players, playerID)
	roomcache[roomKey] = &tRoomCache{
		info:    m,
		players: players,
	}
}

func getRoomCache(roomKey int) *tRoomCache {
	return roomcache[roomKey]
}
