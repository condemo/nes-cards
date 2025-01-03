package store

import (
	"context"

	"github.com/condemo/nes-cards/types"
	"github.com/uptrace/bun"
)

type Store interface {
	CreatePlayer(*types.Player) error
	GetPlayer(int64) (*types.Player, error)
	CreateGame(*types.Game) error
	GetLastGame() (*types.Game, error)
	GetGameList() ([]types.Game, error)
}

type Storage struct {
	db *bun.DB
}

func NewStorage(db *bun.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreatePlayer(p *types.Player) error {
	_, err := s.db.NewInsert().Model(p).
		Returning("*").Exec(context.Background())
	return err
}

func (s *Storage) GetPlayer(id int64) (*types.Player, error) {
	p := new(types.Player)
	err := s.db.NewSelect().Model(p).
		Where("id = ?", id).Limit(1).Scan(context.Background())

	return p, err
}

// TODO: Ineficiente, dos querys en lugar de una
func (s *Storage) CreateGame(g *types.Game) error {
	_, err := s.db.NewInsert().Model(g).
		Returning("*").Exec(context.Background())
	if err != nil {
		return err
	}

	err = s.db.NewSelect().Model(g).
		Relation("Player1").Where("p1 = player1.id").
		Relation("Player2").Where("p2 = player2.id").
		Where("g.id = ?", g.ID).
		Scan(context.Background())

	return err
}

func (s *Storage) GetLastGame() (*types.Game, error) {
	g := new(types.Game)
	err := s.db.NewSelect().Model(g).
		Relation("Player1").Where("p1 = player1.id").
		Relation("Player2").Where("p2 = player2.id").
		Order("g.created_at DESC").Limit(1).
		Scan(context.Background())

	return g, err
}

func (s *Storage) GetGameList() ([]types.Game, error) {
	var pl []types.Game

	err := s.db.NewSelect().Model(&pl).
		Relation("Player1").Where("p1=player1.id").
		Relation("Player2").Where("p2=Player2.id").
		Scan(context.Background())
	if err != nil {
		return nil, err
	}

	return pl, nil
}
