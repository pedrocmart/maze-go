package repository

import (
	"database/sql"

	"github.com/pedrocmart/maze-go/repository/models"
)

type Level interface {
	Create(*models.Level) (*models.Level, error)
	FindByLevelId(int64) ([]*models.Level, error)
	FindAll() ([]*models.Level, error)
}

type LevelRepository struct {
	db *sql.DB
}

func NewLevelRepository(db *sql.DB) *LevelRepository {
	return &LevelRepository{db}
}

func (r *LevelRepository) Create(level *models.Level) (*models.Level, error) {
	var err error

	const queryStr = `INSERT INTO levels (maps) VALUES ($1) RETURNING id`

	tx, err := r.db.Begin()
	if err != nil {
		// TODO add logging
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	err = r.db.QueryRow(queryStr, level.Maps).Scan(&level.Id)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return level, nil
}

func (r *LevelRepository) FindByLevelId(levelId int64) ([]*models.Level, error) {
	const queryStr = `SELECT id,
		maps,
		created_at
	FROM levels
	WHERE id = $1
	ORDER BY created_at DESC`

	params := []interface{}{levelId}

	q, err := r.db.Query(queryStr, params...)
	if err != nil {
		return nil, err
	}

	defer func() {
		q.Close()
	}()

	var levels []*models.Level
	for q.Next() {
		level := models.Level{}

		err = q.Scan(&level.Id, &level.Maps, &level.CreatedAt)
		if err != nil {
			return nil, err
		}

		levels = append(levels, &level)
	}

	if err := q.Err(); err != nil {
		return nil, err
	}

	return levels, nil
}

func (r *LevelRepository) FindAll() ([]*models.Level, error) {
	const stmt = `SELECT id,
		maps, 
		created_at
	FROM levels
	ORDER BY created_at DESC`

	q, err := r.db.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer func() {
		q.Close()
	}()

	var rounds []*models.Level
	for q.Next() {
		round := models.Level{}

		err = q.Scan(&round.Id, &round.Maps, &round.CreatedAt)
		if err != nil {
			return nil, err
		}

		rounds = append(rounds, &round)
	}

	if err := q.Err(); err != nil {
		return nil, err
	}

	return rounds, nil
}
