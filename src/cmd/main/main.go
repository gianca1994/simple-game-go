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
			Username: "gianca",
			Password: "test",
			Level: 1,
			Experience: 0,
			ExperienceToNextLevel: 100,
			DamageMax: 10,
			DamageMin: 5,
			Health: 100,
			DefenseMax: 10,
			DefenseMin: 5,
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
	db.First(&user1, "username = ?", "gianca")
	db.First(&user2, "username = ?", "lucho")

	service.CombatSystem(user1, user2)
}
