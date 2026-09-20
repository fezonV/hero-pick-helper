package models

type OpenDotaMatchup struct {
	HeroID      int64 `json:"hero_id"`
	GamesPlayed int64 `json:"games_played"`
	Wins        int64 `json:"wins"`
}

// считает винрейт исходного героя против HeroID из этого matchup
func (mu OpenDotaMatchup) Winrate() (float32, error) {
	if mu.GamesPlayed == 0 {
		return 0.0, ErrZeroMatches
	}
	return float32(mu.Wins) / float32(mu.GamesPlayed), nil
}
