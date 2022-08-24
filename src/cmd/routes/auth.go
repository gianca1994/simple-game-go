package routes

import (
	"encoding/json"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"simple-game-golang/src/model"
)

var tokenAuth *jwtauth.JWTAuth

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := DbConnection()

	var user model.User

	json.NewDecoder(r.Body).Decode(&user)
	passwordReq := user.Password

	db.Where("username = ?", user.Username).Find(&user)

	if CheckPasswordHash(user.Password, passwordReq) {
		tokenAuth = jwtauth.New("HS512", []byte("secret"), nil)
		_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"id": user.ID})

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"token": "` + tokenString + `"}`))
	} else {
		w.Write([]byte("Incorrect password"))
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
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
	db.Create(&user)

	w.Write([]byte("User created successfully"))
}
