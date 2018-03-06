package co

import (
	"fmt"
)

type ByteSlice []byte

func (b ByteSlice) Len() int {
	return len(b)
}

func (b ByteSlice) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b ByteSlice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByteSlice) String() string {
	return string(b)
}

func (b ByteSlice) Ints() []int {
	ret := make([]int, 0, len(b))
	for _, card := range b {
		ret = append(ret, int(card))
	}
	return ret
}

func (b ByteSlice) Hexs() string {
	if len(b) == 0 {
		return "[]"
	}

	ret := "["
	for _, card := range b {
		ret += fmt.Sprintf("0x%.2x,", int(card))
	}

	return ret[:len(ret)-1] + "]"
}

func ByteSliceToString(val []ByteSlice) string {
	s := ""
	for _, v := range val {
		s += string(v)
	}
	return s
}

func Ints2ByteSlice(val []int) ByteSlice {
	ret := make(ByteSlice, 0, len(val))
	for _, card := range val {
		ret = append(ret, byte(card))
	}
	return ret
}

func (b ByteSlice) Exclude(exclude ByteSlice) ByteSlice {
	dst := make(ByteSlice, 0, len(b))
	exl := make(ByteSlice, 0, len(exclude))
	exl = append(exl, exclude...)
	for _, s := range b {
		fIgnore := false
		for idx, e := range exl {
			if s == e {
				fIgnore = true
				exl[idx] = 0
				break
			}
		}

		if !fIgnore {
			dst = append(dst, s)
		}
	}

	return dst
}

func (b ByteSlice) Sum() byte {
	var sum byte
	for _, v := range b {
		sum += v
	}
	return sum
}

func (b ByteSlice) Index(c byte) int {
	for idx, v := range b {
		if v == c {
			return idx
		}
	}

	return -1
}
