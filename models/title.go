package models

import "database/sql"

type Title struct {
	Tconst         string         `json:"id"`
	TitleType      string         `json:"title_type"`
	PrimaryTitle   string         `json:"primary_title"`
	OriginalTitle  string         `json:"original_title"`
	IsAdult        bool           `json:"is_adult"`
	StartYear      sql.NullInt64  `json:"start_year"`
	EndYear        sql.NullInt64  `json:"end_year"`
	RuntimeMinutes sql.NullInt64  `json:"runtime_minutes"`
	Genres         sql.NullString `json:"genres"`
}
