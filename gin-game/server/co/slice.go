package co

func CombCards(cards ByteSlice, m int, emit func(ByteSlice)) {
	s := make(ByteSlice, m)
	last := m - 1
	var rc func(int, int)
	rc = func(i, n int) {
		for j := n; j < len(cards); j++ {
			s[i] = cards[j]
			// fmt.Println(i, last)
			if i == int(last) {
				emit(s)
			} else {
				rc(i+1, j+1)
			}
		}
		return
	}
	rc(0, 0)
}

func IndexOf(arr []int, v int) int {
	for idx, val := range arr {
		if v == val {
			return idx
		}
	}

	return -1
}

func Exclude(arr []byte, exi []int) []byte {
	ret := make([]byte, 0, len(arr)-len(exi))
	for i, val := range arr {
		if IndexOf(exi, i) == -1 {
			ret = append(ret, val)
		}
	}
	return ret
}
