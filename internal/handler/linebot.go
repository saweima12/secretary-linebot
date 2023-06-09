package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/saweima12/secretary-bot/internal/repository"
	"github.com/saweima12/secretary-bot/internal/secretary"
	"github.com/saweima12/secretary-bot/internal/service"
)

const (
	RECORDED_MSG = "Message recorded"
)

func (r *router) registerLineBotHandler(parent *gin.RouterGroup) {

	// create group.
	group := parent.Group("/bot")

	// select database & colection
	db := r.App.Mongo.Database(secretary.DB_SECRETARY)
	msgRepo := repository.NewBotMessageRepository(db.Collection(secretary.COLLECTION_BOT_MSG))
	botService := service.NewLineBotService(r.App.Bot, msgRepo)

	handler := &LineBotHandler{
		botService: botService,
	}

	// register router.
	group.POST("/update", handler.newUpdate)
	group.GET("/me", func(ctx *gin.Context) {
		ctx.String(200, "ok")
	})

}

type LineBotHandler struct {
	botService service.LineBotService
}

func (h *LineBotHandler) newUpdate(ctx *gin.Context) {
	events, err := h.botService.ParseRequest(ctx.Request)

	if err != nil {
		log.Fatalf("Can't parse request: %+v ", ctx.Request)
	}

	for _, event := range events {
		switch message := event.Message.(type) {
		case *linebot.TextMessage:
			err = h.botService.SaveMessage(event.Source, message)
			if err != nil {
				log.Fatalf("Can't save message: %+v", message)
			}

			err = h.botService.SendMessage(event.Source.UserID, RECORDED_MSG)
			if err != nil {
				log.Fatalf("Can't send message: %+v", message)
			}
		}
	}

}
