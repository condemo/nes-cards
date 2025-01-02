package handlers

import (
	"net/http"

	"github.com/condemo/nes-cards/public/views/core"
	"github.com/condemo/nes-cards/types"
)

type viewsHandler struct{}

func NewViewsHandler() *viewsHandler {
	return &viewsHandler{}
}

func (h *viewsHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /", h.homeView)
	r.HandleFunc("GET /current-game", h.gameView)
	r.HandleFunc("GET /history", h.historyView)
}

func (h *viewsHandler) homeView(w http.ResponseWriter, r *http.Request) {
	RenderTempl(w, r, core.Home())
}

func (h *viewsHandler) gameView(w http.ResponseWriter, r *http.Request) {
	// TODO:
	p1 := types.NewPlayer("Demo1", 80)
	p2 := types.NewPlayer("Demo2", 80)
	t1 := types.NewTower(60)
	t2 := types.NewTower(60)

	RenderTempl(w, r, core.GameView(p1, p2, t1, t2))
}

func (h *viewsHandler) historyView(w http.ResponseWriter, r *http.Request) {
	// TODO:
	RenderTempl(w, r, core.HistoryView())
}
