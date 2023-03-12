package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	modules "github.com/saweima12/secretary-bot/internal/secretary"
)

const (
	BOT_SECRET = "BOT_LINE_SECRET"
	BOT_TOKEN  = "BOT_LINE_TOKEN"
)

func New() *modules.ServerApp {
	// get bot channel secret & token from environment.
	secret := os.Getenv(BOT_SECRET)
	token := os.Getenv(BOT_TOKEN)

	// declare gin server.
	engine := gin.Default()
	bot, err := linebot.New(secret, token)

	if err != nil {
		panic(err)
	}

	// engine.SetTrustedProxies([]string{"172.17.0.1"})

	// create app instance.
	app := &modules.ServerApp{
		Engine: engine,
		Bot:    bot,
	}

	return app
}
