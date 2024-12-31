package types

type Game struct {
	P1 *Player
	P2 *Player
}

func NewGame() *Game {
	g := &Game{}

	return g
}

// LoadPlayers recibe los datos de los Jugadores creados
// y los carga en la partida
func (g *Game) LoadPlayers(p1, p2 Player) error {
	g.P1 = &p1
	g.P2 = &p2

	return nil
}
