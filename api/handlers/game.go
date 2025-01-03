package handlers

import (
	"log"
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

func (h gameHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /new", h.newGameView)
	r.HandleFunc("POST /new", h.newGamePost)
}

func (h gameHandler) newGameView(w http.ResponseWriter, r *http.Request) {
	RenderTempl(w, r, core.NewGameView())
}

func (h gameHandler) newGamePost(w http.ResponseWriter, r *http.Request) {
	var playerHP uint8
	p1 := r.FormValue("player1")
	p2 := r.FormValue("player2")

	pHP := r.FormValue("player-hp")
	if pHP == "" {
		playerHP = 80
	} else {
		p, err := strconv.ParseUint(pHP, 0, 8)
		if err != nil {
			log.Fatal(err)
		}
		playerHP = uint8(p)
	}

	// TODO: Implementar
	// tHP := r.FormValue("tower-hp")
	// if tHP == "" {
	// 	towerHP = 60
	// } else {
	// 	t, err := strconv.ParseUint(pHP, 0, 8)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	towerHP = uint8(t)
	// }

	player1 := types.NewPlayer(p1, playerHP)
	player2 := types.NewPlayer(p2, playerHP)

	if err := h.store.CreatePlayer(player1); err != nil {
		log.Fatal("error creating a player: ", err)
	}
	if err := h.store.CreatePlayer(player2); err != nil {
		log.Fatal("error creating a player: ", err)
	}

	game := types.NewGame(player1.ID, player2.ID)

	if err := h.store.CreateGame(game); err != nil {
		log.Fatal("error creating a game: ", err)
	}

	RenderTempl(w, r, core.GameView(game))
}
