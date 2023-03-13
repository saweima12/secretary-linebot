package service

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/saweima12/secretary-bot/internal/repository"
)

type LineBotService interface {
	ParseRequest(req *http.Request) ([]*linebot.Event, error)
}

type lineBotService struct {
	Client  *linebot.Client
	MsgRepo repository.UserMsgRepository
}

func NewLineBotService(client *linebot.Client, msgRepo repository.UserMsgRepository) LineBotService {
	return &lineBotService{
		Client:  client,
		MsgRepo: msgRepo,
	}
}

func (s *lineBotService) ParseRequest(req *http.Request) ([]*linebot.Event, error) {
	return s.Client.ParseRequest(req)
}

func (s *lineBotService) SaveBotMessage(msg *linebot.Message) error {
	return nil
}
