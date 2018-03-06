package niuniu

import (
	"encoding/gob"
	"gin-game/server/def"
	"os"

	"github.com/inconshreveable/log15"
)

var (
	niuTbl map[string]*def.Niu
	log    = log15.New("module", "loadniuniu")
)

func init() {
	if loadTable() {
		log.Info("牛牛表载入完成")
	} else {
		log.Error("牛牛表载入失败")
		// os.Exit(1)
	}

}

func loadTable() bool {
	file, err := os.OpenFile("gen/niuniu.tbl", os.O_RDONLY, 0644)
	if err != nil {
		log.Error("无法打开牛牛表", "error", err)
		return false
	}

	defer func() {
		file.Close()
	}()

	r := gob.NewDecoder(file)
	err = r.Decode(&niuTbl)
	if err != nil {
		log.Error("无反序列化牛牛表", "error", err)
		return false
	}

	return true
}
