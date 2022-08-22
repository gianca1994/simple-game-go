package model

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	DamageMax  int    `json:"damageMax"`
	DamageMin  int    `json:"damageMin"`
	Health     int    `json:"health"`
	DefenseMax int    `json:"defenseMax"`
	DefenseMin int    `json:"defenseMin"`
}
