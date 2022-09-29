package routes

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	jwt_auth "simple-game-golang/src/internal/jwt_bearer"
	"simple-game-golang/src/model"
)

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
	var userDB model.User
	db.Where("username = ?", user.Username).First(&userDB)

	if CheckPasswordHash(userDB.Password, user.Password) {
		token := jwt_auth.GenerateToken(userDB)

		data, _ := json.Marshal(map[string]string{
			"token": token,
		})

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	} else {
		w.Write([]byte("Invalid username or password"))
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	db := DbConnection()

	var user model.User
	var UserRegisterDTO model.UserRegisterDTO

	json.NewDecoder(r.Body).Decode(&UserRegisterDTO)

	user.Username = UserRegisterDTO.Username
	hash, _ := HashPassword(UserRegisterDTO.Password)
	user.Password = hash


	db.Create(&user)

	w.Write([]byte("User created successfully"))
}


