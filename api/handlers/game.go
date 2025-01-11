package handlers

import (
	"net/http"
	"strconv"

	"github.com/condemo/nes-cards/public/views/core"
	"github.com/condemo/nes-cards/store"
	"github.com/condemo/nes-cards/types"
)

type gameHandler struct {
	store store.Store
}

func NewGameHandler(s store.Store) *gameHandler {
	return &gameHandler{store: s}
}

func (h *gameHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /new", makeHandler(h.newGameView))
	r.HandleFunc("POST /new", makeHandler(h.newGamePost))
}

func (h *gameHandler) newGameView(w http.ResponseWriter, r *http.Request) error {
	gl, err := h.store.GetPlayerList()
	if err != nil {
		return err
	}

	return RenderTempl(w, r, core.NewGameView(gl))
}

func (h *gameHandler) newGamePost(w http.ResponseWriter, r *http.Request) error {
	var game *types.Game
	var playerHP uint8 = 80

	p1 := r.FormValue("player1")
	p2 := r.FormValue("player2")
	pHP := r.FormValue("player-hp")

	// Default Player names if Name is empty
	if p1 == "" {
		p1 = "Player1"
	}
	if p2 == "" {
		p2 = "Player2"
	}

	// Check if pHP is empty
	if pHP != "" {
		h, err := strconv.ParseUint(pHP, 10, 8)
		if err != nil {
			return err
		}
		playerHP = uint8(h)
	}

	// Players create
	player1 := types.NewPlayer(p1, playerHP)
	player2 := types.NewPlayer(p2, playerHP)

	// Check if player already exists
	if ok := h.store.CheckPlayer(player1.Name); !ok {
		if err := h.store.CreatePlayer(player1); err != nil {
			return err
		}
	} else {
		err := h.store.GetPlayerByName(player1)
		if err != nil {
			return err
		}
	}
	if ok := h.store.CheckPlayer(player2.Name); !ok {
		if err := h.store.CreatePlayer(player2); err != nil {
			return err
		}
	} else {
		err := h.store.GetPlayerByName(player2)
		if err != nil {
			return err
		}
	}

	// Create Game
	game = types.NewGame(player1.ID, player2.ID)
	if err := h.store.CreateGame(game); err != nil {
		return err
	}

	return RenderTempl(w, r, core.GameView(game))
}
