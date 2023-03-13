package secretary

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

type LineBotClient interface {
	ParseRequest(r *http.Request) ([]*linebot.Event, error)
}

func Test() {

}
