package handlers

import (
	"net/http"

	"github.com/condemo/nes-cards/public/views/core"
)

type viewsHandler struct{}

func NewViewsHandler() *viewsHandler {
	return &viewsHandler{}
}

func (h *viewsHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /", h.homeView)
	r.HandleFunc("GET /welcome", h.welcomeView)
	r.HandleFunc("GET /game", h.gameView)
	r.HandleFunc("GET /history", h.historyView)
}

func (h *viewsHandler) homeView(w http.ResponseWriter, r *http.Request) {
	RenderTempl(w, r, core.Home())
}

func (h *viewsHandler) welcomeView(w http.ResponseWriter, r *http.Request) {
	// TODO:
	RenderTempl(w, r, core.WelcomeView())
}

func (h *viewsHandler) gameView(w http.ResponseWriter, r *http.Request) {
	// TODO:
	RenderTempl(w, r, core.GameView())
}

func (h *viewsHandler) historyView(w http.ResponseWriter, r *http.Request) {
	// TODO:
	RenderTempl(w, r, core.HistoryView())
}
