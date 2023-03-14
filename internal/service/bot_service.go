package service

import (
	"encoding/json"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/saweima12/secretary-bot/internal/entity"
	"github.com/saweima12/secretary-bot/internal/repository"
)

type LineBotService interface {
	ParseRequest(req *http.Request) ([]*linebot.Event, error)
	SaveMessage(source *linebot.EventSource, msg *linebot.TextMessage) error
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

func (s *lineBotService) SaveMessage(source *linebot.EventSource, msg *linebot.TextMessage) error {
	j, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	model := entity.UserMsg{
		UserId:  source.UserID,
		Message: msg.Text,
		Raw:     string(j),
	}

	_, err = s.MsgRepo.InsertOne(model)
	if err != nil {
		return err
	}

	return nil
}
