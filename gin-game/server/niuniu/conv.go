package niuniu

import (
	"gin-game/server/co"
)

func Names2Card(names []string) co.ByteSlice {
	arr := make(co.ByteSlice, 0, len(names))
	for _, name := range names {
		arr = append(arr, name2Card[name])
	}
	return arr
}

func Cards2Name(cards []byte) string {
	if len(cards) == 0 {
		return "[]"
	}

	ret := "["
	for _, card := range cards {
		ret += card2Name[card] + ","
	}

	return ret[:len(ret)-1] + "]"
}
