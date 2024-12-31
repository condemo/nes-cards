package types

type Player struct {
	Name string
	hp   uint8
}

// NewPlayer recibe un nombre e instancia un nuevo jugador
func NewPlayer(name string, hp uint8) *Player {
	p := &Player{
		Name: name,
		hp:   hp,
	}

	return p
}

// Hit resta `dmg` a la HP del jugador
func (p *Player) Hit(dmg uint8) {
	p.hp = p.hp - dmg
}
