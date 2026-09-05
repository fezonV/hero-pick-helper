package models

type Matchup struct {
	HeroID      int64 `json:"hero_id"`
	GamesPlayed int64 `json:"games_played"`
	Wins        int64 `json:"wins"`
}

// Считает винрейт для каждого персонажа против того, которого мы передали в паттерне
func (mu Matchup) Winrate() float32 {
	return float32(mu.GamesPlayed) / float32(mu.Wins)
}
