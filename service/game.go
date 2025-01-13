package service

import "github.com/condemo/nes-cards/types"

type GameController struct {
	game *types.Game
}

func NewGameController() *GameController {
	return &GameController{}
}

func (gc *GameController) SetGame(g *types.Game) {
	gc.game = g
}

func (gc *GameController) NextTurn() {
	// TODO:
}
