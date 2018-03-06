package niuniu

import (
	"fmt"
	"gin-game/server/co"
	"gin-game/server/def"
	"sort"
)

func QueryScore(cards []string) *def.Niu {
	arr := Names2Card(cards)
	sort.Sort(co.ByteSlice(arr))

	var val *def.Niu
	var exist bool
	if val, exist = niuTbl[string(arr)]; exist {
		fmt.Println("arr:", Cards2Name(arr), "ret:", val)
	}

	return val
}

func QueryScore2(cards co.ByteSlice) *def.Niu {
	sort.Sort(co.ByteSlice(cards))
	var val *def.Niu
	var exist bool
	if val, exist = niuTbl[string(cards)]; exist {
		fmt.Println("cards:", Cards2Name(cards), "NiuIndex:", Cards2Name(val.Index.Niu), "NiuRemain:", Cards2Name(val.Index.Remain), "Niu:", val.Number, "Score:", val.Score)
	}

	return val
}
