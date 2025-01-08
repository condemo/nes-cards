package handlers

import (
	"fmt"
	"net/http"

	"github.com/condemo/nes-cards/public/views/components"
	"github.com/condemo/nes-cards/store"
)

type playerHandler struct {
	store store.Store
}

func NewPlayerHandler(s store.Store) *playerHandler {
	return &playerHandler{store: s}
}

func (h *playerHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /modal", h.getNewPlayerModal)
	r.HandleFunc("POST /", h.createPlayer)
}

func (h *playerHandler) getNewPlayerModal(w http.ResponseWriter, r *http.Request) {
	RenderTempl(w, r, components.NewPlayerModal())
}

func (h *playerHandler) createPlayer(w http.ResponseWriter, r *http.Request) {
	p := r.FormValue("new-name")
	fmt.Println("creating player -> ", p)

	// TODO:
}
