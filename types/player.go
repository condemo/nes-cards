package types

import "github.com/uptrace/bun"

type Player struct {
	bun.BaseModel `bun:"table:players,alias:p"`

	ID   int64  `bun:",pk,autoincrement"`
	Name string `bun:",notnull"`
	HP   uint8  `bun:",nullzero"`
	al   []AlteredEffect
}

// NewPlayer recibe un nombre e instancia un nuevo jugador
func NewPlayer(name string, hp uint8) *Player {
	p := &Player{
		Name: name,
		HP:   hp,
	}

	return p
}

func (p *Player) TakeDMG(dmg uint8) {
	p.HP = p.HP - dmg
}

func (p *Player) AddAlteredEffect(a AlteredEffect) {
	p.al = append(p.al, a)
}

func (p *Player) CleanAltered() {
	if len(p.al) > 0 {
		p.al = p.al[:len(p.al)-1]
	}
}

func (p *Player) ApplyAlteredStack() {
	for _, ae := range p.al {
		ae.Apply()
	}
}
