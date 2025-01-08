package handlers

import (
	"log"
	"net/http"

	"github.com/condemo/nes-cards/public/views/core"
	"github.com/condemo/nes-cards/store"
	"github.com/condemo/nes-cards/types"
)

type viewsHandler struct {
	db store.Store
}

func NewViewsHandler(s store.Store) *viewsHandler {
	return &viewsHandler{
		db: s,
	}
}

func (h *viewsHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /", h.homeView)
	r.HandleFunc("GET /home", h.frontView)
	r.HandleFunc("GET /current-game", h.gameView)
	r.HandleFunc("GET /history", h.historyView)
}

func (h *viewsHandler) homeView(w http.ResponseWriter, r *http.Request) {
	RenderTempl(w, r, core.Home())
}

func (h *viewsHandler) frontView(w http.ResponseWriter, r *http.Request) {
	RenderTempl(w, r, core.FrontView())
}

func (h *viewsHandler) gameView(w http.ResponseWriter, r *http.Request) {
	// TODO:
	var game *types.Game

	game, err := h.db.GetLastGame()
	if err != nil {
		RenderTempl(w, r, core.EmptyView())
		return
	}

	RenderTempl(w, r, core.GameView(game))
}

func (h *viewsHandler) historyView(w http.ResponseWriter, r *http.Request) {
	// TODO:
	var gl []types.Game
	gl, err := h.db.GetGameList()
	if err != nil {
		log.Fatal("historyView db error -> ", err)
	}

	if len(gl) == 0 {
		RenderTempl(w, r, core.HistoryEmpty())
		return
	}

	RenderTempl(w, r, core.HistoryView(gl))
}
