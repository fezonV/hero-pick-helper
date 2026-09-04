package model

type Attribute string

const (
	AttributeAgility   Attribute = "agi"
	AttributeStrength  Attribute = "str"
	AttributeIntellect Attribute = "int"
	AttributeUniversal Attribute = "all"
)

type Attack string

const (
	Ranged Attack = "Ranged"
	Melee  Attack = "Melee"
)

type Role string

const (
	Carry     Role = "Carry"
	Support   Role = "Support"
	Nuker     Role = "Nuker"
	Disabler  Role = "Disabler"
	Jungler   Role = "Jungler"
	Durable   Role = "Durable"
	Escape    Role = "Escape"
	Pusher    Role = "Pusher"
	Initiator Role = "Initiator"
)

type HeroModel struct {
	Id             int64     `json:"id"`
	Name           string    `json:"name"`
	Localized_name string    `json: "localized_name"`
	Primary_attr   Attribute `json:"primary_attr"`
	Attack_type    Attack
	Roles          []Role
}

func CreateHeroModel(
	id int64,
	name string,
	localized_name string,
	primary_attr Attribute,
	attack_type Attack,
	roles []Role,
) (*HeroModel, error) {
	return &HeroModel{
		Id:             id,
		Name:           name,
		Localized_name: localized_name,
		Primary_attr:   primary_attr,
		Attack_type:    attack_type,
		Roles:          roles,
	}, nil
}
