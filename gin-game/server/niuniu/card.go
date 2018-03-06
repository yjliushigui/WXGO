package niuniu

import (
	"gin-game/server/co"
	"math/rand"
	"time"
)

const (
	colorDiamond = 1
	colorClub    = 2
	colorHeart   = 3
	colorSpade   = 4
)

var (
	poker = []byte{
		// 方块
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d,
		// 梅花
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d,
		// 红桃
		0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d,
		// 黑桃
		0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d,
	}

	name2Card = map[string]byte{
		"方块A":  0x11,
		"方块2":  0x12,
		"方块3":  0x13,
		"方块4":  0x14,
		"方块5":  0x15,
		"方块6":  0x16,
		"方块7":  0x17,
		"方块8":  0x18,
		"方块9":  0x19,
		"方块10": 0x1a,
		"方块J":  0x1b,
		"方块Q":  0x1c,
		"方块K":  0x1d,

		"梅花A":  0x21,
		"梅花2":  0x22,
		"梅花3":  0x23,
		"梅花4":  0x24,
		"梅花5":  0x25,
		"梅花6":  0x26,
		"梅花7":  0x27,
		"梅花8":  0x28,
		"梅花9":  0x29,
		"梅花10": 0x2a,
		"梅花J":  0x2b,
		"梅花Q":  0x2c,
		"梅花K":  0x2d,

		"红桃A":  0x31,
		"红桃2":  0x32,
		"红桃3":  0x33,
		"红桃4":  0x34,
		"红桃5":  0x35,
		"红桃6":  0x36,
		"红桃7":  0x37,
		"红桃8":  0x38,
		"红桃9":  0x39,
		"红桃10": 0x3a,
		"红桃J":  0x3b,
		"红桃Q":  0x3c,
		"红桃K":  0x3d,

		"黑桃A":  0x41,
		"黑桃2":  0x42,
		"黑桃3":  0x43,
		"黑桃4":  0x44,
		"黑桃5":  0x45,
		"黑桃6":  0x46,
		"黑桃7":  0x47,
		"黑桃8":  0x48,
		"黑桃9":  0x49,
		"黑桃10": 0x4a,
		"黑桃J":  0x4b,
		"黑桃Q":  0x4c,
		"黑桃K":  0x4d,
	}
	card2Name = map[byte]string{
		0x11: "方块A",
		0x12: "方块2",
		0x13: "方块3",
		0x14: "方块4",
		0x15: "方块5",
		0x16: "方块6",
		0x17: "方块7",
		0x18: "方块8",
		0x19: "方块9",
		0x1a: "方块10",
		0x1b: "方块J",
		0x1c: "方块Q",
		0x1d: "方块K",

		0x21: "梅花A",
		0x22: "梅花2",
		0x23: "梅花3",
		0x24: "梅花4",
		0x25: "梅花5",
		0x26: "梅花6",
		0x27: "梅花7",
		0x28: "梅花8",
		0x29: "梅花9",
		0x2a: "梅花10",
		0x2b: "梅花J",
		0x2c: "梅花Q",
		0x2d: "梅花K",

		0x31: "红桃A",
		0x32: "红桃2",
		0x33: "红桃3",
		0x34: "红桃4",
		0x35: "红桃5",
		0x36: "红桃6",
		0x37: "红桃7",
		0x38: "红桃8",
		0x39: "红桃9",
		0x3a: "红桃10",
		0x3b: "红桃J",
		0x3c: "红桃Q",
		0x3d: "红桃K",

		0x41: "黑桃A",
		0x42: "黑桃2",
		0x43: "黑桃3",
		0x44: "黑桃4",
		0x45: "黑桃5",
		0x46: "黑桃6",
		0x47: "黑桃7",
		0x48: "黑桃8",
		0x49: "黑桃9",
		0x4a: "黑桃10",
		0x4b: "黑桃J",
		0x4c: "黑桃Q",
		0x4d: "黑桃K",
	}

	theCards co.ByteSlice
)

const (
	// SuitCount 花色数
	SuitCount = 4
	// SuitPointCount 每种花色的牌数 1~10
	SuitPointCount = 13
	// SameCardCount 相同牌数量
	SameCardCount = 1
	// TotalCount 套牌总数
	TotalCount = SuitCount * SuitPointCount * SameCardCount
)

// SysCard 游戏牌，由系统掌握（非玩家可以操作）
type SysCard struct {
	workCards co.ByteSlice
	workPos   int
}

// 构建一副牌（只需要调用一次）
func init() {
	theCards = make(co.ByteSlice, 0, TotalCount)
	for n := 0; n < SameCardCount; n++ {
		for card := range card2Name {
			theCards = append(theCards, card)
		}
	}
}

func getSysCards() co.ByteSlice {
	cards := make(co.ByteSlice, 0, len(card2Name))
	for card := range card2Name {
		cards = append(cards, card)
	}
	return cards
}

// NewSysCard 系统管理的卡牌
func NewSysCard() *SysCard {
	return &SysCard{
		workCards: make(co.ByteSlice, TotalCount, TotalCount),
	}
}

// Init 初始化
func (c *SysCard) Init() {
	rand.Seed(time.Now().Unix())
}

// Shuffle 洗牌
func (c *SysCard) Shuffle() {
	indexs := rand.Perm(len(theCards))
	for i, randIndex := range indexs {
		c.workCards[i] = theCards[randIndex]
	}
	c.workPos = 0
}

// Deal 发牌
func (c *SysCard) Deal(count int) co.ByteSlice {
	if count <= 0 {
		return co.ByteSlice{}
	}

	simDeal := c.workPos + count

	if simDeal > len(theCards) {
		log.Debug("发牌数量超出一副牌总数")
		return nil
	}

	cards := c.workCards[c.workPos:simDeal]
	c.workPos = simDeal
	return cards
}

// Remain 剩余牌数
func (c *SysCard) Remain() int {
	return len(c.workCards) - c.workPos
}

// GetCards 获取套牌(不重复)
func GetCards() co.ByteSlice {
	cards := make(co.ByteSlice, 0, len(card2Name))
	for card := range card2Name {
		cards = append(cards, card)
	}
	return cards
}

func (c *SysCard) ShuffleEx(cards co.ByteSlice) {
	for idx, card := range cards {
		c.workCards[idx] = card
	}

	cardCount := make(map[byte]int)
	for _, card := range cards {
		cardCount[card]++
	}

	n := 0
	indexs := rand.Perm(len(theCards))
	for _, randIndex := range indexs {
		card := theCards[randIndex]
		cardCount[card]--
		if cardCount[card] < 0 {
			if n+len(cards) < len(theCards) {
				c.workCards[n+len(cards)] = card
				n++
			}
		}
	}
	c.workPos = 0
}
