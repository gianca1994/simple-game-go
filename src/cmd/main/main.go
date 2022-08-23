package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"simple-game-golang/src/internal/database"
	"simple-game-golang/src/model"
	"simple-game-golang/src/service"
)

func main() {

	db := database.NewPostgreSQL()

	if db == nil {
		fmt.Println("Error connecting to database")
		os.Exit(0)
	}

	// CREATE USER:
	/*
		db.Create(&model.User{
			Name:       "gianca",
			DamageMax:  10,
			DamageMin:  5,
			Health:     50,
			DefenseMax: 3,
			DefenseMin: 1,
		})
	*/

	// Find users:
	/*
		var users []model.User
		db.Find(&users)
	*/

	// Find user by id:
	/*
		var user model.User
		db.First(&user, 1)
		fmt.Println(user)
	*/

	var user1, user2 model.User

	db.First(&user1, "name = ?", "Gianca")
	db.First(&user2, "name = ?", "lucho")

	service.CombatSystem(user1, user2)
}
