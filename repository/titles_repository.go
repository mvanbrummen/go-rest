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

func (r *TitlesRepository) FetchTitle(id string) (*models.Title, error) {
	t := new(models.Title)

	query := `select tconst, titleType, primaryTitle, 
	originalTitle, isAdult, startYear, endYear, 
	runtimeMinutes, genres
	from imdb.titles
	where tconst = $1`

	row := r.db.QueryRow(query, id)

	err := row.Scan(
		&t.Tconst,
		&t.TitleType,
		&t.PrimaryTitle,
		&t.OriginalTitle,
		&t.IsAdult,
		&t.StartYear,
		&t.EndYear,
		&t.RuntimeMinutes,
		&t.Genres,
	)

	if err != nil {
		panic(err)
	}

	return t, nil
}
