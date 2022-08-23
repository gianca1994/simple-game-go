package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"simple-game-golang/src/internal/database"
	"simple-game-golang/src/model"

	"github.com/go-chi/chi/v5"
)

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	db := database.NewPostgreSQL()

	if db == nil {
		fmt.Println("Error connecting to database")
		os.Exit(0)
	}

	var users []model.User
	db.Find(&users)
	data, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func getUserByName(w http.ResponseWriter, r *http.Request) {
	db := database.NewPostgreSQL()

	if db == nil {
		fmt.Println("Error connecting to database")
		os.Exit(0)
	}

	name := chi.URLParam(r, "name")
	var user model.User
	db.Where("username = ?", name).First(&user)
	data, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {

	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Get("/", getAllUsers)
		r.Get("/{name}", getUserByName)
	})

	http.ListenAndServe(":8080", r)

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

	// var user1, user2 model.User
	// db.First(&user1, "username = ?", "gianca")
	// db.First(&user2, "username = ?", "lucho")

	// service.CombatSystem(user1, user2)
}
