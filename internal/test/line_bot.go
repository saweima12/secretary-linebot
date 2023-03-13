package test

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/saweima12/secretary-bot/internal/secretary"
)

type mockLineBot struct{}

func NewMockLineBot() secretary.LineBotClient {
	return &mockLineBot{}
}

func (c *mockLineBot) ParseRequest(r *http.Request) ([]*linebot.Event, error) {

	return nil, nil
}
