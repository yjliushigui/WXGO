package game

import (
	"container/list"
	"gin-game/server/co"
	"sort"
)

// PlayerCard 玩家拥有牌
type PlayerCard struct {
	// 关联的玩家ID
	playerID int

	// 玩家拥有的牌
	cards *list.List
}

// NewPlayerCard 玩家牌
func NewPlayerCard() *PlayerCard {
	return &PlayerCard{
		cards: list.New(),
	}
}

// Init 初始化
func (c *PlayerCard) Init(playerID int, cards co.ByteSlice) {
	c.Reset()
	c.playerID = playerID
	sort.Sort(cards)
	for _, card := range cards {
		c.cards.PushBack(card)
	}
}

// Reset 重置
func (c *PlayerCard) Reset() {}

// 获取手牌
func (c PlayerCard) getHandCards() []int {
	arr := make([]int, 0, c.cards.Len())
	for e := c.cards.Front(); e != nil; e = e.Next() {
		arr = append(arr, e.Value.(int))
	}
	return arr
}
