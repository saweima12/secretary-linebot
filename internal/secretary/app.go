package secretary

import (
	"context"
	"log"
	"time"

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
	// Create database connection.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := app.Mongo.Connect(ctx)
	if err != nil {
		return err
	}

	// check db connection is ok.
	err = app.Mongo.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return err
	}

	// Start gin server.
	err = app.Engine.Run(port)

	if err != nil {
		return err
	}

	return nil
}
