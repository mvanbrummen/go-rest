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

func (r *TitlesRepository) SearchByTitle(title string, limit int) ([]*models.Title, error) {

	query := `select tconst, titleType, primaryTitle, 
	originalTitle, isAdult, startYear, endYear, 
	runtimeMinutes, genres
	from imdb.titles
	where primaryTitle like '%' || $1 || '%' 
	limit $2`

	rows, err := r.db.Query(query, title, limit)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := make([]*models.Title, 0)
	for rows.Next() {
		t := new(models.TitleDAO)

		err = rows.Scan(
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
		result = append(result, t.ToTitle())
	}

	return result, nil
}

func (r *TitlesRepository) FetchTitle(id string) (*models.Title, error) {
	t := new(models.TitleDAO)

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
		return nil, nil
	}

	return t.ToTitle(), nil
}
