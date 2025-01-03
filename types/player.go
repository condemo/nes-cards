package types

import "github.com/uptrace/bun"

type Player struct {
	bun.BaseModel `bun:"table:players,alias:p"`

	ID   int64  `bun:",pk,autoincrement"`
	Name string `bun:",notnull"`
	HP   uint8  `bun:",notnull,nullzero"`
}

// NewPlayer recibe un nombre e instancia un nuevo jugador
func NewPlayer(name string, hp uint8) *Player {
	p := &Player{
		Name: name,
		HP:   hp,
	}

	return p
}
