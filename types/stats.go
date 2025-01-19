package types

import "github.com/uptrace/bun"

type Stats struct {
	bun.BaseModel `bun:"table:stats,alias:t"`

	ID       int64 `bun:",pk,autoincrement"`
	GameID   int64 `bun:",notnull"`
	PlayerID int64 `bun:",notnull"`
	HP       uint8 `bun:",nullzero" validator:"gte=0,lte=255"`
	T1HP     uint8 `bun:",nullzero"`
	T2HP     uint8 `bun:",nullzero"`
}

// NewStats recibe GameID y PlayerID además de una cantidad de HP e instancia dos Torres
func NewStats(gid, pid int64, php, thp uint8) *Stats {
	// TODO:

	s := &Stats{
		GameID:   gid,
		PlayerID: pid,
		HP:       php,
		T1HP:     thp,
		T2HP:     thp,
	}

	return s
}

func (t *Stats) Validate() error {
	err := validate.Struct(t)
	return err
}
