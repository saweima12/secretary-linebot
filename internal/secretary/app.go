package secretary

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServerApp struct {
	Engine *gin.Engine
	Bot    *linebot.Client
	Mongo  *mongo.Client
}

func (app *ServerApp) Run(port string) error {

	// Start gin server.
	err := app.Engine.Run(port)

	if err != nil {
		return err
	}

	return nil
}
