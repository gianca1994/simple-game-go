package service

import (
	"fmt"
	. "github.com/ymzuiku/hit"
	"simple-game-golang/src/model"
)

func CombatSystem(attacker, defender model.User) {
	roundCounter := 0
	var attackerDMG, defenderDMG = 0, 0

	for attacker.Health > 0 && defender.Health > 0 {
		roundCounter++
		fmt.Printf("Round %d\n", roundCounter)
		fmt.Printf(attacker.Username+" [Health: %d]\n", attacker.Health)
		fmt.Printf(defender.Username+" [Health: %d\n", defender.Health)

		defender.Health, attackerDMG = Attack(attacker, defender)
		if defender.Health <= 0 {
			defenderDMG = 0
			break
		}

		attacker.Health, defenderDMG = Attack(defender, attacker)
		if attacker.Health <= 0 {
			attackerDMG = 0
			break
		}

		fmt.Printf("Gianca: [Damage Caused: %d]\n", attackerDMG)
		fmt.Printf("Lucho: [Damage Caused: %d]\n", defenderDMG)
	}

	fmt.Printf("Gianca: [Health: %d, Damage Caused: %d]\n", attacker.Health, attackerDMG)
	fmt.Printf("Lucho: [Health: %d, Damage Caused: %d]\n", defender.Health, defenderDMG)

	winner := If(attacker.Health > 0, attacker.Username, defender.Username)
	fmt.Println("The winner is: ", winner)
}
