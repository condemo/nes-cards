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
	r.HandleFunc("GET /modal", makeHandler(h.getNewPlayerModal))
	r.HandleFunc("POST /", makeHandler(h.createPlayer))
}

func (h *playerHandler) getNewPlayerModal(w http.ResponseWriter, r *http.Request) error {
	return RenderTempl(w, r, components.NewPlayerModal())
}

func (h *playerHandler) createPlayer(w http.ResponseWriter, r *http.Request) error {
	p := r.FormValue("new-name")

	if p == "" {
		w.WriteHeader(http.StatusBadRequest)
		return RenderTempl(w, r, components.NewPlayerError("Player Name is Empty"))
	}

	if ok := h.store.CheckPlayer(p); ok {
		w.WriteHeader(http.StatusConflict)
		return RenderTempl(w, r, components.NewPlayerError(fmt.Sprintf("%s already exists", p)))
	}

	player := types.NewPlayer(p, 80)
	if err := h.store.CreatePlayer(player); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return RenderTempl(w, r, components.NewPlayerError("Something went wrong call Gus"))
	}

	pl, err := h.store.GetPlayerList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return RenderTempl(w, r, components.NewPlayerError("Something went wrong call Gus"))
	}

	w.WriteHeader(http.StatusCreated)
	return RenderTempl(w, r, core.NewGameView(pl))
}
