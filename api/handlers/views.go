package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/condemo/nes-cards/public/views/core"
	"github.com/condemo/nes-cards/store"
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
	r.HandleFunc("GET /current-game", h.gameView)
	r.HandleFunc("GET /history", h.historyView)
}

func (h *viewsHandler) homeView(w http.ResponseWriter, r *http.Request) {
	RenderTempl(w, r, core.Home())
}

func (h *viewsHandler) gameView(w http.ResponseWriter, r *http.Request) {
	// TODO:

	game, err := h.db.GetLastGame()
	if err != nil {
		log.Fatal(err)
	}

	RenderTempl(w, r, core.GameView(game))
}

func (h *viewsHandler) historyView(w http.ResponseWriter, r *http.Request) {
	// TODO:
	gl, err := h.db.GetGameList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", gl)
	RenderTempl(w, r, core.HistoryView())
}
