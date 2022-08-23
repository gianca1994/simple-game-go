package service

import (
	"math/rand"
	"simple-game-golang/src/model"
	"time"
)

func getDefense(user model.User) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(user.DefenseMax-user.DefenseMin+1) + user.DefenseMin
}

func getDamage(user model.User) int {
	var damage = 0
	rand.Seed(time.Now().UnixNano())
	damage = rand.Intn(user.DamageMax-user.DamageMin+1) + user.DamageMin
	return damage
}

func Attack(attacker, defender model.User) (int, int) {
	damage := getDamage(attacker)
	//defenseDefender := getDefense(defender)
	defender.Health -= damage
	return defender.Health, damage
}
