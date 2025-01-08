package handlers

import (
	"fmt"
	"net/http"

	"github.com/condemo/nes-cards/public/views/components"
	"github.com/condemo/nes-cards/public/views/core"
	"github.com/condemo/nes-cards/store"
	"github.com/condemo/nes-cards/types"
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

	if p == "" {
		w.WriteHeader(http.StatusBadRequest)
		RenderTempl(w, r, components.NewPlayerError("Player Name is Empty"))
		return
	}

	if ok := h.store.CheckPlayer(p); ok {
		w.WriteHeader(http.StatusConflict)
		RenderTempl(w, r, components.NewPlayerError(fmt.Sprintf("%s already exists", p)))
		return
	}

	player := types.NewPlayer(p, 80)
	if err := h.store.CreatePlayer(player); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		RenderTempl(w, r, components.NewPlayerError("Something went wrong call Gus"))
		return
	}

	pl, err := h.store.GetPlayerList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		RenderTempl(w, r, components.NewPlayerError("Something went wrong call Gus"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	RenderTempl(w, r, core.NewGameView(pl))
}
