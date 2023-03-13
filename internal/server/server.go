package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/saweima12/secretary-bot/internal/handler"
	"github.com/saweima12/secretary-bot/internal/secretary"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ENV_BOT_SECRET = "BOT_LINE_SECRET"
	ENV_BOT_TOKEN  = "BOT_LINE_TOKEN"
	ENV_MONGO_URI  = "BOT_MONGO_URI"
)

func New() *secretary.ServerApp {
	// get bot channel botSecret & token from environment.
	botSecret := os.Getenv(ENV_BOT_SECRET)
	botToken := os.Getenv(ENV_BOT_TOKEN)

	if botSecret == "" || botToken == "" {
		panic("Please set environment variable BOT_LINE_SECRET and BOT_LINE_TOKEN")
	}

	// declare gin server.
	engine := gin.Default()

	// intialize Linebot client.
	bot, err := linebot.New(botSecret, botToken)
	if err != nil {
		panic(err)
	}

	// get mongodb uri from environment variable.
	dbUri := os.Getenv(ENV_MONGO_URI)

	// create mongodb Client.
	dbClient, err := initMongoDB(dbUri)
	if err != nil {
		panic(err)
	}

	// create app instance.
	app := &secretary.ServerApp{
		Engine: engine,
		Bot:    bot,
		Mongo:  dbClient,
	}

	// register router.
	router := handler.NewRouter(app)
	router.Init()

	return app
}

func initMongoDB(dbUri string) (*mongo.Client, error) {
	// declare options & ctx

	log.Printf("TEST DBURI: %s", dbUri)
	opts := options.Client().ApplyURI(dbUri)
	client, err := mongo.NewClient(opts)

	if err != nil {
		return nil, err
	}

	return client, nil
}
