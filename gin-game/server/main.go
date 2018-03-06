package main

import (
	"gin-game/server/co"
	"gin-game/server/game"
	"gin-game/server/msg"
	"time"

	"go.uber.org/zap"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

var (
	handCards co.ByteSlice
)

func serve() {
	game.Init()

	r := gin.Default()
	logger, _ := zap.NewProduction()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	r.GET("/ws", func(c *gin.Context) {
		msg.WSHandler(c.Writer, c.Request)
	})

	r.Static("/css", "statics/css")
	r.Static("/image", "statics/image")
	r.LoadHTMLFiles("statics/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.Run(":15000")
}

func main() {
	serve()
}
