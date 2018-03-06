package co

import (
	"container/list"
	"encoding/json"
	"reflect"

	"github.com/name5566/leaf/log"
)

// TypeName 获取类型名称
func TypeName(v interface{}) string {
	return reflect.TypeOf(v).Elem().Name()
}

func TypeNames(lst *list.List) string {
	s := "["
	for e := lst.Front(); e != nil; e = e.Next() {
		s += TypeName(e.Value) + ", "
	}
	s = s[:len(s)-2] + "]"
	return s
}

func ConvertMsg(data []byte, m interface{}) bool {
	log.Debug("ConvertMsg %v", string(data))
	err := json.Unmarshal(data, m)
	if err != nil {
		log.Error("ConvertMsg %v", err)
		return false
	}
	return true
}
