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
	var towerHP uint8
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

	tHP := r.FormValue("tower-hp")
	if tHP == "" {
		towerHP = 60
	} else {
		t, err := strconv.ParseUint(pHP, 0, 8)
		if err != nil {
			log.Fatal(err)
		}
		towerHP = uint8(t)
	}

	player1 := types.NewPlayer(p1, uint8(playerHP))
	player2 := types.NewPlayer(p2, uint8(playerHP))
	tower1 := types.NewTower(uint8(towerHP))
	tower2 := types.NewTower(uint8(towerHP))

	game := types.NewGame()
	game.LoadPlayers(*player1, *player2)

	if err := h.store.CreateGame(game); err != nil {
		log.Fatal("error creating a game: ", err)
	}

	RenderTempl(w, r, core.GameView(player1, player2, tower1, tower2))
}
