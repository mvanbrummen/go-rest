package repository

import (
	"log"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"tconst", "titleType", "primaryTitle", "originalTitle",
		"isAdult", "startYear", "endYear", "runtimeMinutes", "genres"}).
		AddRow("tc00001", "movie", "Labryinth", "Labryinth", true, 1985, 1985, 108, "action")

	query := "select tconst, titleType, primaryTitle, originalTitle, isAdult, startYear, endYear, runtimeMinutes, genres from imdb.titles where tconst = \\$1"

	mock.ExpectQuery(query).WillReturnRows(rows)

	repo := NewTitlesRepository(db)

	title, err := repo.FetchTitle("tc00001")

	assert.NoError(t, err)
	assert.NotNil(t, title)
}

func TestSearchByTitle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"tconst", "titleType", "primaryTitle", "originalTitle",
		"isAdult", "startYear", "endYear", "runtimeMinutes", "genres"}).
		AddRow("tc00001", "movie", "Labryinth", "Labryinth", true, 1985, 1985, 108, "action").
		AddRow("tc00002", "tv", "Labryinth", "Labryinth", true, 2003, 2003, 30, "documentary")

	query := `select tconst, titleType, primaryTitle, 
	originalTitle, isAdult, startYear, endYear, 
	runtimeMinutes, genres
	from imdb.titles
	where primaryTitle like '%' || \\$1 || '%' 
	limit \\$2`

	mock.ExpectQuery(query).WillReturnRows(rows)

	repo := NewTitlesRepository(db)

	titles, err := repo.SearchByTitle("Labryinth", 2)

	assert.NoError(t, err)
	assert.Len(t, titles, 2)
}
