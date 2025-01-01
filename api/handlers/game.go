package handlers

import (
	"fmt"
	"net/http"

	"github.com/condemo/nes-cards/public/views/core"
)

type gameHandler struct{}

func NewGameHandler() *gameHandler {
	return &gameHandler{}
}

func (h gameHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /new", h.newGameView)
	r.HandleFunc("POST /new", h.newGamePost)
}

func (h gameHandler) newGameView(w http.ResponseWriter, r *http.Request) {
	RenderTempl(w, r, core.NewGameView())
}

func (h gameHandler) newGamePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post made")
	RenderTempl(w, r, core.GameView())
}
