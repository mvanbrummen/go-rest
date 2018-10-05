package repository

import (
	"database/sql"
)

type TitlesRepository struct {
	db *sql.DB
}

func NewTitlesRepository(db *sql.DB) *TitlesRepository {
	return &TitlesRepository{db}
}

// func (*TitlesRepository) FetchTitle(id string) (models.Title, error) {
// 	return nil, nil
// }
