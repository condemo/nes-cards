package types

import (
	"context"
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

var _ bun.BeforeAppendModelHook = (*Game)(nil)

// BeforeAppendModel gets the current time in GMT+1 and set it in CreatedAt field
func (g *Game) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		location, err := time.LoadLocation("Europe/Madrid")
		if err != nil {
			return err
		}
		// FIX: Añado una hora de free por el cambio de hora, debería ser automático
		g.CreatedAt = time.Now().In(location).Add(time.Hour)
	}

	return nil
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
