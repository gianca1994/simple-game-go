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

func addUser(w http.ResponseWriter, r *http.Request) {
	db := database.NewPostgreSQL()

	if db == nil {
		fmt.Println("Error connecting to database")
		os.Exit(0)
	}

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	user.Level = 1
	user.Experience = 0
	user.ExperienceToNextLevel = 100
	user.DamageMax = 10
	user.DamageMin = 5
	user.Health = 100
	user.DefenseMax = 10
	user.DefenseMin = 5
	db.Save(&user)

	data, _ := json.Marshal(user)
	w.Write(data)
}

func main() {

	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Get("/", getAllUsers)
		r.Get("/{name}", getUserByName)
		r.Post("/", addUser)
	})

	http.ListenAndServe(":8080", r)


	// var user1, user2 model.User
	// db.First(&user1, "username = ?", "gianca")
	// db.First(&user2, "username = ?", "lucho")

	// service.CombatSystem(user1, user2)
}
