package def

import (
	"gin-game/server/co"
)

// 游戏状态
const (
	// GameNotStart 游戏未开始
	GameNotStart = 0
	// GameReady 游戏准备OK
	GameReady = 1
	// GamePlaying 游戏正在耍
	GamePlaying = 2
	// GameOver 游戏结束了，总结算后
	GameOver = 3
)

type ISysCard interface {
	Init()
	Shuffle()
	Deal(count int) co.ByteSlice
	Remain() int
	// 测试洗牌接口
	ShuffleEx(cards co.ByteSlice)
}

type IGame interface {
	Start()
	Run(playerID int, data []byte)
	Stop()
	State() int
	// Score(playerID int) int
	// Recovery(playerID int) *msg.S2C_Recovery
	// SendFinalBalance(playerID int)
	// BroadcastFinalBalance()
}
