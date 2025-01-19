package types

import (
	"github.com/uptrace/bun"
)

type Player struct {
	bun.BaseModel `bun:"table:players,alias:p"`

	ID   int64  `bun:",pk,autoincrement"`
	Name string `bun:",notnull" validate:"required,min=3,max=10,alphanum"`
}

// NewPlayer recibe un nombre e instancia un nuevo jugador
func NewPlayer(name string) *Player {
	p := &Player{
		Name: name,
	}

	return p
}

func (p *Player) Validate() error {
	err := validate.Struct(p)
	return err
}
