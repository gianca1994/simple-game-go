package main

import (
	"crud-golang-sqlite/src/model"
	"crud-golang-sqlite/src/service"
)

func main() {
	gianca := model.User{
		Id:         1,
		Name:       "Gianca",
		DamageMax:  10,
		DamageMin:  5,
		Health:     50,
		DefenseMax: 3,
		DefenseMin: 1,
	}
	lucho := model.User{
		Id:         2,
		Name:       "Lucho",
		DamageMax:  10,
		DamageMin:  5,
		Health:     50,
		DefenseMax: 3,
		DefenseMin: 1,
	}
	service.CombatSystem(gianca, lucho)
}
