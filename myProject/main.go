package main

import (
	_ "myProject/routers"
	"github.com/astaxie/beego"
	"time"
	"strconv"
)

func main() {
	beego.AddFuncMap("strtotime", strtotime)
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

func strtotime(in string) (out string) {
	timestamp, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		timestamp = 0
	}
	tm := time.Unix(timestamp, 10)
	out = tm.Format("2006-01-02 03:04:05")
	return
}
