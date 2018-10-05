package repository

import (
	"database/sql"

	"github.com/mvanbrummen/go-rest/models"
)

type TitlesRepository struct {
	db *sql.DB
}

func NewTitlesRepository(db *sql.DB) *TitlesRepository {
	return &TitlesRepository{db}
}

func (*TitlesRepository) FetchTitle(id string) (*models.Title, error) {
	return &models.Title{
		"tconst00000",
		"movie",
		"Die Hard",
		"Die Hard",
		true,
		1988,
		1988,
		130,
		[]string{"action", "romance"},
	}, nil
}
