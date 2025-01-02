package types

import (
	"time"

	"github.com/uptrace/bun"
)

type Game struct {
	bun.BaseModel `bun:"table:games,alias:g"`
	ID            int64     `bun:",pk,autoincrement"`
	P1            string    `bun:"player1,notnull"`
	P2            string    `bun:"player2,notnull"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func NewGame() *Game {
	g := &Game{}

	return g
}

// LoadPlayers recibe los datos de los Jugadores creados
// y los carga en la partida
func (g *Game) LoadPlayers(p1, p2 Player) error {
	g.P1 = p1.Name
	g.P2 = p2.Name

	return nil
}
