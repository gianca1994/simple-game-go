package service

import (
	"crud-golang-sqlite/src/model"
	"fmt"
	. "github.com/ymzuiku/hit"
)

func CombatSystem(attacker model.User, defender model.User) {
	roundCounter := 0

	for attacker.Health > 0 && defender.Health > 0 {
		roundCounter++
		fmt.Printf("Round %d\n", roundCounter)

		defender.Health = Attack(attacker, defender)
		if defender.Health <= 0 {
			break
		}

		attacker.Health = Attack(defender, attacker)
		if attacker.Health <= 0 {
			break
		}

		fmt.Println("Gianca: ", attacker)
		fmt.Println("Lucho: ", defender)
	}

	fmt.Println("Gianca: ", attacker)
	fmt.Println("Lucho: ", defender)

	winner := If(attacker.Health > 0, attacker.Name, defender.Name)
	fmt.Println("The winner is: ", winner)
}
