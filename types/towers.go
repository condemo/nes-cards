package types

type Tower struct {
	hp uint8
}

// NewTower recibe una cantidad de HP e instancia una Torre
func NewTower(hp uint8) *Tower {
	// TODO:
	t := &Tower{hp: hp}

	return t
}

func (t *Tower) GetHP() uint8 {
	return t.hp
}
