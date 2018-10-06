package models

import (
	"database/sql"
	"strings"
)

type TitleDAO struct {
	Tconst         string
	TitleType      string
	PrimaryTitle   string
	OriginalTitle  string
	IsAdult        bool
	StartYear      sql.NullInt64
	EndYear        sql.NullInt64
	RuntimeMinutes sql.NullInt64
	Genres         sql.NullString
}

type Title struct {
	ID             string   `json:"id"`
	TitleType      string   `json:"title_type"`
	PrimaryTitle   string   `json:"primary_title"`
	OriginalTitle  string   `json:"original_title"`
	IsAdult        bool     `json:"is_adult"`
	StartYear      *int64   `json:"start_year"`
	EndYear        *int64   `json:"end_year"`
	RuntimeMinutes *int64   `json:"runtime_minutes"`
	Genres         []string `json:"genres"`
}

func (t *TitleDAO) ToTitle() *Title {
	var startYear *int64 = nil
	if t.StartYear.Valid {
		startYear = &t.StartYear.Int64
	}

	var endYear *int64 = nil
	if t.EndYear.Valid {
		endYear = &t.EndYear.Int64
	}

	var runtimeMinutes *int64 = nil
	if t.RuntimeMinutes.Valid {
		runtimeMinutes = &t.RuntimeMinutes.Int64
	}

	var genres []string
	if t.Genres.Valid {
		genres = strings.Split(t.Genres.String, ",")
	}

	return &Title{
		strings.TrimSpace(t.Tconst),
		t.TitleType,
		t.PrimaryTitle,
		t.OriginalTitle,
		t.IsAdult,
		startYear,
		endYear,
		runtimeMinutes,
		genres,
	}
}
