package handlers

import (
	"errors"
	"fmt"
	"net/http"

	apiErrors "github.com/condemo/nes-cards/api/api_errors"
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
		return apiErrors.NewApiError(
			errors.New("player name is empty"),
			"Player Name is Empty",
			http.StatusBadRequest,
		)
	}

	if ok := h.store.CheckPlayer(p); ok {
		return apiErrors.NewApiError(
			errors.New("player already exists"),
			fmt.Sprintf("%s already exists", p),
			http.StatusConflict,
		)
	}

	player := types.NewPlayer(p, 80)
	if err := h.store.CreatePlayer(player); err != nil {
		return err
	}

	pl, err := h.store.GetPlayerList()
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return RenderTempl(w, r, core.NewGameView(pl))
}
