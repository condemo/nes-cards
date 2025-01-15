package store

import (
	"context"

	"github.com/condemo/nes-cards/types"
	"github.com/uptrace/bun"
)

type Store interface {
	CreatePlayer(*types.Player) error
	CheckPlayer(string) bool
	GetPlayerById(*types.Player) error
	GetPlayerByName(*types.Player) error
	GetPlayerList() ([]types.Player, error)
	CreateTowerList([]*types.Tower) error
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

func (s *Storage) CheckPlayer(name string) bool {
	if err := s.db.NewSelect().
		Model(&types.Player{}).
		Where("name = ?", name).
		Scan(context.Background()); err != nil {
		return false
	}

	return true
}

func (s *Storage) GetPlayerById(p *types.Player) error {
	err := s.db.NewSelect().Model(p).
		Where("id = ?", p.ID).Scan(context.Background())

	return err
}

func (s *Storage) GetPlayerByName(p *types.Player) error {
	err := s.db.NewSelect().Model(p).
		Where("name = ?", p.Name).Scan(context.Background())

	return err
}

func (s *Storage) GetPlayerList() ([]types.Player, error) {
	var pl []types.Player
	err := s.db.NewSelect().Model(&pl).
		Order("id ASC").Limit(20).
		Scan(context.Background())

	return pl, err
}

func (s *Storage) CreateTowerList(tl []*types.Tower) error {
	_, err := s.db.NewInsert().Model(&tl).
		Returning("*").Exec(context.Background())
	return err
}

func (s *Storage) CreateGame(g *types.Game) error {
	// TODO: Ineficiente, dos querys en lugar de una
	_, err := s.db.NewInsert().Model(g).
		Returning("*").Exec(context.Background())
	if err != nil {
		return err
	}

	err = s.db.NewSelect().Model(g).
		Relation("Player1").Where("p1id = player1.id").
		Relation("Player2").Where("p2id = player2.id").
		Relation("Towers1", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.ExcludeColumn("*")
		}).
		Relation("Towers2", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.ExcludeColumn("*")
		}).
		Where("g.id = ?", g.ID).
		Scan(context.Background())

	return err
}

func (s *Storage) GetLastGame() (*types.Game, error) {
	g := new(types.Game)
	err := s.db.NewSelect().Model(g).
		Relation("Player1").Where("p1id = player1.id").
		Relation("Player2").Where("p2id = player2.id").
		Relation("Towers1", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.ExcludeColumn("*")
		}).
		Relation("Towers2", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.ExcludeColumn("*")
		}).
		Order("g.created_at DESC").Limit(1).
		Scan(context.Background())

	return g, err
}

func (s *Storage) GetGameList() ([]types.Game, error) {
	var pl []types.Game

	err := s.db.NewSelect().Model(&pl).
		Relation("Player1").Where("p1id=player1.id").
		Relation("Player2").Where("p2id=Player2.id").
		Relation("Towers1", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.ExcludeColumn("*")
		}).
		Relation("Towers2", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.ExcludeColumn("*")
		}).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}

	return pl, nil
}
