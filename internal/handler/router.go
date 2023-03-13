package handler

import "github.com/saweima12/secretary-bot/internal/secretary"

type Router interface {
	Init()
}

type router struct {
	App *secretary.ServerApp
}

func NewRouter(app *secretary.ServerApp) Router {
	return &router{
		App: app,
	}
}

func (r *router) Init() {

	router := r.App.Engine
	group := router.Group("/api/v1")
	// register bot handler.
	r.registerLineBotHandler(group)
}
