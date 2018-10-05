package models

import "database/sql"

type Title struct {
	Tconst         string        `json:"id"`
	TitleType      string        `json:"title_type"`
	PrimaryTitle   string        `json:"primary_title"`
	OriginalTitle  string        `json:"original_title"`
	IsAdult        bool          `json:"is_adult"`
	StartYear      int           `json:"start_year"`
	EndYear        sql.NullInt64 `json:"end_year"`
	RuntimeMinutes int           `json:"runtime_minutes"`
	Genres         string        `json:"genres"`
}
