package store

import (
	"context"

	"github.com/condemo/nes-cards/types"
	"github.com/uptrace/bun"
)

type Store interface {
	CreateGame(*types.Game) error
	GetGameList() ([]*types.Game, error)
}

type Storage struct {
	db *bun.DB
}

func NewStorage(db *bun.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateGame(g *types.Game) error {
	_, err := s.db.NewInsert().Model(g).
		Returning("*").Exec(context.Background())
	return err
}

func (s *Storage) GetGameList() ([]*types.Game, error) {
	return nil, nil
}
