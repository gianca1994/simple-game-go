package routes

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"os"
	"simple-game-golang/src/internal/database"
	"simple-game-golang/src/model"
)



func DbConnection() *gorm.DB {
	db := database.NewPostgreSQL()

	if db == nil {
		fmt.Println("Error connecting to database")
		os.Exit(0)
	}
	return db
}


func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := DbConnection()

	var users []model.User
	db.Find(&users)
	data, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func GetUserByName(w http.ResponseWriter, r *http.Request) {
	db := DbConnection()

	name := chi.URLParam(r, "name")
	var user model.User
	db.Where("username = ?", name).First(&user)
	data, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

