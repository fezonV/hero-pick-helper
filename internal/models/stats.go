package models

type Stats struct {
	Id             int64     `json:"id"`
	Name           string    `json:"name"`
	Localized_name string    `json:"localized_name"`
	Primary_attr   Attribute `json:"primary_attr"`
	Attack_type    Attack    `json:"attack_type"`
	Roles          []Role    `json:"roles"`

	ProPick int64 `json:"pro_pick"` // прошники
	ProWin  int64 `json:"pro_win"`
	ProBan  int64 `json:"pro_ban"`

	Pick1 int64 `json:"1_pick"` // рекруты
	Win1  int64 `json:"1_win"`

	Pick2 int64 `json:"2_pick"` // стражи
	Win2  int64 `json:"2_win"`

	Pick3 int64 `json:"3_pick"` // рыцари
	Win3  int64 `json:"3_win"`

	Pick4 int64 `json:"4_pick"` // герои
	Win4  int64 `json:"4_win"`

	Pick5 int64 `json:"5_pick"` // легенды
	Win5  int64 `json:"5_win"`

	Pick6 int64 `json:"6_pick"` // властелины
	Win6  int64 `json:"6_win"`

	Pick7 int64 `json:"7_pick"` // божества
	Win7  int64 `json:"7_win"`
}

func (s Stats) Winrate(rank int) (float32, error) {
	if rank > 7 || rank <= 0 {
		return 0.0, ErrBadRank
	}
	switch rank {
	case 1:
		if s.Pick1 == 0 {
			return 0.0, ErrZeroMatches
		}
		return float32(s.Win1) / float32(s.Pick1), nil
	case 2:
		if s.Pick2 == 0 {
			return 0.0, ErrZeroMatches
		}
		return float32(s.Win2) / float32(s.Pick2), nil
	case 3:
		if s.Pick3 == 0 {
			return 0.0, ErrZeroMatches
		}
		return float32(s.Win3) / float32(s.Pick3), nil
	case 4:
		if s.Pick4 == 0 {
			return 0.0, ErrZeroMatches
		}
		return float32(s.Win4) / float32(s.Pick4), nil
	case 5:
		if s.Pick5 == 0 {
			return 0.0, ErrZeroMatches
		}
		return float32(s.Win5) / float32(s.Pick5), nil
	case 6:
		if s.Pick6 == 0 {
			return 0.0, ErrZeroMatches
		}
		return float32(s.Win6) / float32(s.Pick6), nil
	case 7:
		if s.Pick7 == 0 {
			return 0.0, ErrZeroMatches
		}
		return float32(s.Win7) / float32(s.Pick7), nil
	default:
		return 0.0, ErrBadRank
	}
}
