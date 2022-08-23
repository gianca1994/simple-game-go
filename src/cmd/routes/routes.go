package routes

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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

func AddUser(w http.ResponseWriter, r *http.Request) {
	db := DbConnection()

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	hash, _ := HashPassword(user.Password)

	user.Password = hash
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
