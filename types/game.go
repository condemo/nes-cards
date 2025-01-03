package types

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Game struct {
	bun.BaseModel `bun:"table:games,alias:g"`

	ID        int64 `bun:",pk,autoincrement"`
	P1        int64
	P2        int64
	Player1   Player    `bun:"rel:belongs-to,join:p1=id"`
	Player2   Player    `bun:"rel:belongs-to,join:p2=id"`
	Winner    string    `bun:"winner"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
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

func NewGame(pid1, pid2 int64) *Game {
	g := &Game{
		P1: pid1,
		P2: pid2,
	}

	return g
}
