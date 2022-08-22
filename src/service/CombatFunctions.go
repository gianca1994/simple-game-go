package service

import (
	"crud-golang-sqlite/src/model"
	"math/rand"
	"time"
)

func getDamageAndDefense(user model.User) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(user.DamageMax - user.DamageMin + 1) + user.DamageMin
}

func Attack(attacker, defender model.User) int {
	damage := getDamageAndDefense(attacker)
	defenseDefender := getDamageAndDefense(defender)

	defender.Health -= damage / defenseDefender
	return defender.Health
}
