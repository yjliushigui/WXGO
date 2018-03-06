package main

import (
	"encoding/gob"
	"gin-game/server/def"
	"gin-game/server/co"
	"gin-game/server/niuniu/niuname"
	"os"
	"sort"

	"github.com/inconshreveable/log15"
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

	log = log15.New("module", "gen_niuniu_tbl")
)

// genTable 生存牛牛表
func genTable() {
	arrNiu := make(map[string]*def.Niu)
	co.CombCards(co.ByteSlice(poker), 5, func(c5 co.ByteSlice) {
		niu := calcNiuNum(c5)
		niu.Score = calcCardScore(niu.Number, c5)
		sort.Sort(c5)
		arrNiu[string(c5)] = niu
	})

	log.Info("组合生成完成", "牛牛牌组合数", len(arrNiu))

	file, err := os.OpenFile("niuniu.tbl", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0640)
	if err != nil {
		log.Info("创建 da2.data 文件失败:", "error", err)
		return
	}

	defer file.Close()
	w := gob.NewEncoder(file)
	err = w.Encode(arrNiu)
	if err != nil {
		log.Info("牛牛组合序列化失败！", "error", err)
		return
	}
}

func getColor(card byte) byte {
	return card / 0x10
}

func getPoint(card byte) byte {
	return card % 0x10
}

func getNiuPoint(card byte) byte {
	point := card % 0x10
	if point > 9 {
		return 10
	}

	return point
}

func calcCardScore(number int, cards []byte) int {
	score := number * 1000
	for _, card := range cards {
		score += int(getPoint(card))*10 + int(getColor(card))/0x10
	}
	return score
}

func calcNiuNum(cards []byte) *def.Niu {
	// -- 返回: {牛数, {3张牛牌}, {剩余2张}}
	ret := &def.Niu{Number: -1}

	var tmp byte
	// -- 最大牛在cards中的索引
	maxNiuIndex := make([]byte, 3, 3)
	// -- 余下的2张
	var remain2Cards []byte
	// -- 花色相同(同花)
	thArray := make([]byte, 5, 5)
	// -- 点数相同(炸弹)
	zdArray := make([]byte, 5, 5)
	// -- 牛牛规则下的点数
	points := make([]byte, 0, 5)
	// -- 三张牌点数求和结果
	var sum3CardsPoint byte
	// -- 卡牌(详细)
	var cardI byte
	var cardN byte
	var cardK byte

	// -- 顺花牛点数
	hs1 := []byte{2, 3, 4, 5, 6}
	hs2 := []byte{4, 5, 6, 7, 8}

	for i := 0; i < len(cards); i++ {
		cardI = cards[i]
		sum3CardsPoint = sum3CardsPoint + getNiuPoint(cardI)
		for n := i + 1; n < len(cards); n++ {
			cardN = cards[n]
			if getColor(cardI) == getColor(cardN) {
				thArray[i] = thArray[i] + 1
			}

			if getColor(cardI) == getColor(cardN) {
				zdArray[i] = zdArray[i] + 1
			}

			for k := i + 1; k < len(cards); k++ {
				cardK = cards[k]
				tmp = (getNiuPoint(cardI) + getNiuPoint(cardN) + getNiuPoint(cardK)) % 10
				if tmp == 0 {
					remain2Cards = co.Exclude(cards, []int{i, n, k})
					// -- 统计剩余2张牌的点数
					for _, v := range remain2Cards {
						tmp = tmp + getNiuPoint(v)

					}

					// -- 记住最大牛数
					if int(tmp) > ret.Number {
						ret.Number = int(tmp)
						// -- 记住该3张组合
						maxNiuIndex[0] = byte(i)
						maxNiuIndex[1] = byte(n)
						maxNiuIndex[2] = byte(k)
						ret = &def.Niu{Index: def.NiuIndex{
							Niu:    []byte{cards[i], cards[n], cards[k]},
							Remain: remain2Cards,
						}}

					}
				}
			}
		}
	}

	// -- 修改没牛与牛牛点数
	if ret.Number != -1 {
		if ret.Number == 0 {
			ret.Number = niuname.NiuNiu
		}
	} else {
		ret.Number = 0
	}

	// -- 判断是否是银花牛
	if sum3CardsPoint == byte(len(cards)*10) {
		// -- 判断是否是金牛
		ret.Number = niuname.JinHua
		for _, v := range cards {
			if getPoint(v) == 10 {
				ret.Number = niuname.YinHua
				break
			}
		}
	}

	// -- 五小牛
	tmp = 0
	var less5 = true
	for _, v := range cards {
		if getPoint(v) >= 5 {
			less5 = false
			break
		}
		tmp += getPoint(v)
	}

	if less5 && tmp <= 10 {
		ret.Number = niuname.WuXiao
	}

	// -- 判断炸弹
	sort.SliceStable(zdArray, func(i, j int) bool { return zdArray[i] < zdArray[j] })
	// -- 判断花炸
	if zdArray[len(zdArray)-1] == 3 {
		if ret.Number == niuname.JinHua {
			ret.Number = niuname.HuaZha
		} else {
			ret.Number = niuname.ZhaDan
		}
	}

	// -- 判断同花
	sort.SliceStable(points, func(i, j int) bool { return points[i] < points[j] })
	sort.SliceStable(thArray, func(i, j int) bool { return thArray[i] < thArray[j] })
	if thArray[len(thArray)-1] == 4 {
		c1 := true
		c2 := true
		for i, v := range points {
			if c1 && hs1[i] != v {
				c1 = false
			}

			if c2 && hs2[i] != v {
				c2 = false
			}
		}

		if c1 || c2 {
			ret.Number = niuname.HuaShun
		} else {
			ret.Number = niuname.TongHua
		}
	}

	if ret.Number == niuname.None {
		ret.Index.Niu = []byte{}
		ret.Index.Remain = []byte{}
	}

	return ret
}

func main() {
	genTable()

	// arr := niuniu.Names2Card(strings.Split("方块A,梅花3,梅花7,黑桃2,梅花J", ","))
	// sort.Sort(arr)
	// fmt.Println("test", arr.Hexs())
	// ret := calcNiuNum(arr)
	// fmt.Println("ret", ret)
}
