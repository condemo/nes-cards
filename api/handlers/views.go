package handlers

import (
	"fmt"
	"net/http"

	"github.com/condemo/nes-cards/public/views/core"
	"github.com/condemo/nes-cards/service"
	"github.com/condemo/nes-cards/store"
	"github.com/condemo/nes-cards/types"
)

type viewsHandler struct {
	db store.Store
	gc *service.GameController
}

func NewViewsHandler(s store.Store, gc *service.GameController) *viewsHandler {
	return &viewsHandler{
		db: s,
		gc: gc,
	}
}

func (h *viewsHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /", makeHandler(h.homeView))
	r.HandleFunc("GET /home", makeHandler(h.frontView))
	r.HandleFunc("GET /current-game", makeHandler(h.gameView))
	r.HandleFunc("GET /history", makeHandler(h.historyView))
	r.HandleFunc("GET /error", makeHandler(h.getErrorView))
}

func (h *viewsHandler) homeView(w http.ResponseWriter, r *http.Request) error {
	return RenderTempl(w, r, core.Home())
}

func (h *viewsHandler) frontView(w http.ResponseWriter, r *http.Request) error {
	return RenderTempl(w, r, core.FrontView())
}

func (h *viewsHandler) gameView(w http.ResponseWriter, r *http.Request) error {
	var game *types.Game

	game, err := h.db.GetLastGame()
	if err != nil {
		return RenderTempl(w, r, core.EmptyView())
	}
	h.gc.SetGame(game)

	sl, err := h.db.GetGameStats(game.ID)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", sl)

	return RenderTempl(w, r, core.GameView(sl))
}

func (h *viewsHandler) historyView(w http.ResponseWriter, r *http.Request) error {
	// TODO:
	var gl []types.Game
	gl, err := h.db.GetGameList()
	if err != nil {
		return err
	}

	if len(gl) == 0 {
		return RenderTempl(w, r, core.HistoryEmpty())
	}

	return RenderTempl(w, r, core.HistoryView(gl))
}

func (h *viewsHandler) getErrorView(w http.ResponseWriter, r *http.Request) error {
	return RenderTempl(w, r, core.InternalErrorView())
}
