package secretary

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type ServerApp struct {
	Engine *gin.Engine
	Bot    *linebot.Client
}

func (app *ServerApp) Run(port string) error {
	err := app.Engine.Run(port)
	if err != nil {
		return err
	}
}
