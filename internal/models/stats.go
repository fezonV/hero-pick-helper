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
	Win4  int64

	Pick5 int64 `json:"5_pick"` // легенды
	Win5  int64

	Pick6 int64 `json:"6_pick"` // властелины
	Win6  int64

	Pick7 int64 `json:"7_pick"` // божества
	Win7  int64

	Pick8 int64 `json:"8_pick"` // титаны
	Win8  int64
}
