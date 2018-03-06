package game

import (
	"container/list"
	"container/ring"
	"fmt"
	"gin-game/server/co"
	"gin-game/server/def"
	"gin-game/server/niuniu"

	"reflect"

	"github.com/name5566/leaf/log"
)

type tPlayerRound struct {
	// 押注倍数
	bet int
}

type tRoundCache struct {
	players []tPlayerRound
}

type Game struct {
	// 参与游戏的玩家ID
	players []int
	// 房间号，控制游戏规则
	roomKey int
	// 系统牌
	syscard def.ISysCard
	// 玩家牌
	playerCard map[int]*PlayerCard

	// 玩家顺序
	order *ring.Ring
	// 每局缓存
	round *tRoundCache
	// 局数
	nRound int
	// 庄家
	dealer int
	// 游戏状态
	state int

	// 缓存玩家请求
	rquestCache map[int]interface{}
	// 等待指定玩家ID
	responsed map[int]bool
	// 流程处理与验证消息类型是否对应 key = flow handle, value = msg type
	flow2MsgTypeMap map[string]interface{}

	flowHandler func(int, []byte)

	// 游戏流程
	flow *ring.Ring
}

// New 构建游戏对象
func New(roomKey int, players []int) *Game {
	order := ring.New(len(players))
	playerCard := make(map[int]*PlayerCard)
	for _, playerIDn := range players {
		playerCard[playerIDn] = NewPlayerCard()
		order.Value = playerIDn
		order = order.Next()
	}

	game := &Game{
		players:    players,
		order:      order,
		syscard:    niuniu.NewSysCard(),
		playerCard: playerCard,
	}

	return game
}

// Start 初始化
func (c *Game) Start() {
	c.round = &tRoundCache{}
	c.state = def.GameReady

	// 增加局数
	c.nRound++
	// 定庄：第一局房主当庄
	if c.dealer == 0 {
		c.orderTo(c.players[0])
		c.dealer = c.players[0]
	}

	// 开始游戏流程
	c.startFlow(c.players, c.handleFirstDraw)
}

// startFlow 流程跳转并执行
func (c *Game) startFlow(playerIDs []int, flowHandler func(int, []byte)) {
	c.flowTo(flowHandler)
	c.makeRequstCache(playerIDs)
}

// orderTo 操作玩家跳转
func (c *Game) orderTo(playerID int) {
	var find bool
	for i := 0; i < c.order.Len(); i++ {
		if c.order.Value.(int) == playerID {
			find = true
			break
		}

		c.order = c.order.Next()
	}

	if !find {
		panic(fmt.Sprintf("orderTo(%d)", playerID))
	}
}

// makeRequstCache 建立缓存
func (c *Game) makeRequstCache(playerIDs []int) {
	c.responsed = make(map[int]bool)
	c.rquestCache = make(map[int]interface{})
	for idx := range playerIDs {
		c.rquestCache[playerIDs[idx]] = nil
	}
}

// flowTo 流程跳转
func (c *Game) flowTo(flowHandler interface{}) {
	fFind := false
	for n := 0; n < c.flow.Len(); n++ {
		if flowKey(c.flow.Value) == flowKey(flowHandler) {
			fFind = true
			break
		}
		c.flow = c.flow.Next()
	}
	if !fFind {
		panic("no flow")
	}
	c.flowHandler = c.flow.Value.(func(int, []byte))
}

func flowKey(handler interface{}) string {
	return fmt.Sprintf("%v", handler)
}

// Run 运行游戏流程
func (c *Game) Run(playerID int, m []byte) {
	// 校验游戏是否能够处理消息类型
	msgType, ok := c.flow2MsgTypeMap[flowKey(c.flow.Value)]
	if !ok {
		//log.Debug("player(%d): recv not register msg type (%v/%v))", playerID, msgType, reflect.TypeOf(m))
		return
	}

	if c.state != def.GamePlaying && c.state != def.GameReady {
		log.Debug("room(%d) game over, cannot process %s", c.roomKey, co.TypeName(m))
		return
	}

	// 验证当游戏流程当前步骤是否能处理该类型消息
	if msgType != reflect.TypeOf(m) {
		subType := false
		if reflect.TypeOf(msgType) == reflect.TypeOf(list.New()) {
			for e := msgType.(*list.List).Front(); e != nil; e = e.Next() {
				if e.Value == reflect.TypeOf(m) {
					subType = true
					break
				}

			}
			if !subType {
				log.Debug("player(%d): current flow(%v) cannot process msg type %v", playerID, co.TypeNames(msgType.(*list.List)), reflect.TypeOf(m))
				return
			}
		} else {
			log.Debug("player(%d): current flow(%v) cannot process msg type %v", playerID, msgType, reflect.TypeOf(m))
			return
		}
	}

	_, exist := c.rquestCache[playerID]
	if !exist {
		log.Debug("player(%d): not room player %v, %v", playerID, c.rquestCache, co.TypeName(m))
		return
	}

	// 执行处理
	c.exec(playerID, m)
}

func (c *Game) exec(playerID int, data []byte) {
	c.flowHandler = c.flow.Value.(func(int, []byte))
	c.flowHandler(playerID, data)
}
