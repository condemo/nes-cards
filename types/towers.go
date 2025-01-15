package types

import "github.com/uptrace/bun"

type Tower struct {
	bun.BaseModel `bun:"table:towers,alias:t"`

	ID       int64 `bun:",pk,autoincrement"`
	GameID   int64 `bun:",notnull"`
	PlayerID int64 `bun:",notnull"`
	HP       uint8 `bun:",nullzero"`
}

// NewTower recibe GameID y PlayerID además de una cantidad de HP e instancia dos Torres
func NewTower(gid, pid int64, hp uint8) (*Tower, *Tower) {
	// TODO:
	t := &Tower{
		GameID:   gid,
		PlayerID: pid,
		HP:       hp,
	}
	t2 := *t

	return t, &t2
}
