package models

type StratzMatchup struct {
	HeroID      int64   `json:"heroId2"`
	GamesPlayed int     `json:"matchCount"`
	Wins        int     `json:"winCount"`
	Winrate     float64 `json:"winsAverage"`
}
