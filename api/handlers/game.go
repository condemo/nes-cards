package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	apiErrors "github.com/condemo/nes-cards/api/api_errors"
	"github.com/condemo/nes-cards/public/views/core"
	"github.com/condemo/nes-cards/service"
	"github.com/condemo/nes-cards/store"
	"github.com/condemo/nes-cards/types"
)

type gameHandler struct {
	store store.Store
	gc    *service.GameController
}

func NewGameHandler(s store.Store, gc *service.GameController) *gameHandler {
	return &gameHandler{store: s, gc: gc}
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
	var towerHP uint8 = 60

	p1 := r.FormValue("player1")
	p2 := r.FormValue("player2")
	pHP := r.FormValue("player-hp")
	tHP := r.FormValue("tower-hp")

	// Default Player names if Name is empty
	if p1 == "" {
		p1 = "Player1"
	}
	if p2 == "" {
		p2 = "Player2"
	}

	// Check if pHP is not empty
	if pHP != "" {
		h, err := strconv.ParseUint(pHP, 10, 8)
		if err != nil {
			return apiErrors.NewApiError(
				err, "hp must be an integer less than 255", http.StatusBadRequest)
		}
		playerHP = uint8(h)
	}

	// Check if tHP is not empty
	if tHP != "" {
		h, err := strconv.ParseUint(tHP, 10, 8)
		if err != nil {
			return apiErrors.NewApiError(
				err, "tower hp must be an integer less than 255", http.StatusBadRequest)
		}
		towerHP = uint8(h)
	}

	// Players create
	player1 := types.NewPlayer(p1, playerHP)
	player2 := types.NewPlayer(p2, playerHP)

	if ok := h.store.CheckPlayer(player1.Name); !ok {
		if err := h.store.CreatePlayer(player1); err != nil {
			return err
		}
	} else {
		if err := h.store.GetPlayerByName(player1); err != nil {
			return err
		}
		fmt.Printf("%+v", player1)
		player1.HP = playerHP
		fmt.Printf("%+v", player1)
	}
	if ok := h.store.CheckPlayer(player2.Name); !ok {
		if err := h.store.CreatePlayer(player2); err != nil {
			return err
		}
	} else {
		if err := h.store.GetPlayerByName(player2); err != nil {
			return err
		}
		player2.HP = playerHP
	}

	// Create Game
	game = types.NewGame(player1.ID, player2.ID)
	if err := h.store.CreateGame(game); err != nil {
		return err
	}

	p1t1, p1t2 := types.NewTower(game.ID, player1.ID, towerHP)
	p2t1, p2t2 := types.NewTower(game.ID, player2.ID, towerHP)

	tl := []*types.Tower{p1t1, p1t2, p2t1, p2t2}
	for _, t := range tl {
		if err := t.Validate(); err != nil {
			return apiErrors.NewApiError(
				err,
				"towers hp must be more than 0 and less than 255",
				http.StatusBadRequest)
		}
	}
	if err := h.store.CreateTowerList(tl); err != nil {
		return err
	}

	game, err := h.store.GetLastGame()
	if err != nil {
		return err
	}

	h.gc.SetGame(game)

	return RenderTempl(w, r, core.GameView(game))
}
