package types

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Game struct {
	bun.BaseModel `bun:"table:games,alias:g"`

	ID        int64 `bun:",pk,autoincrement"`
	P1ID      int64
	P2ID      int64
	Player1   Player    `bun:"rel:belongs-to,join:p1id=id"`
	Player2   Player    `bun:"rel:belongs-to,join:p2id=id"`
	Towers1   []*Tower  `bun:"rel:has-many,join:p1id=player_id,join:id=game_id"`
	Towers2   []*Tower  `bun:"rel:has-many,join:p2id=player_id,join:id=game_id"`
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
		P1ID:   pid1,
		P2ID:   pid2,
		Winner: "none",
	}

	return g
}
