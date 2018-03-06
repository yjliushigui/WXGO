package co

import (
	"container/list"
)

func ListToSlice(lst *list.List) ByteSlice {
	array := make(ByteSlice, 0, lst.Len())
	for e := lst.Front(); e != nil; e = e.Next() {
		array = append(array, e.Value.(byte))
	}
	return array
}
