package routes

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"gorm.io/gorm"
	"net/http"
	"os"
	"simple-game-golang/src/internal/database"
	jwt_auth "simple-game-golang/src/internal/jwt_bearer"
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



func Home(w http.ResponseWriter, r *http.Request) {
	port := ":8080"
	available_routes, _ := json.Marshal(map[string]string{
		"POST, Login": "http://localhost" + port + "/auth/login",
		"POST, Register": "http://localhost" + port + "/auth/register",
		"GET, Profile": "http://localhost" + port + "/profile",
		"GET, All Users": "http://localhost" + port + "/users",
		"GET, By Username": "http://localhost" + port + "/users/{username}",
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(available_routes)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	claims := jwt_auth.TokenGetClaims(jwtauth.TokenFromHeader(r))

	if claims == nil {
		w.Write([]byte("Invalid token"))
		return
	}

	if claims["role"] != "admin" {
		w.Write([]byte("You don't have permission to access this resource"))
		return
	}

	db := DbConnection()
	var users []model.User
	var usersResponseDto []model.UserResponseDTO

	db.Find(&users)

	// TODO: Optimize this code
	for _, user := range users {
		usersResponseDto = append(usersResponseDto, model.UserToUserResponseDTO(user))
	}

	data, _ := json.Marshal(usersResponseDto)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func GetUserByName(w http.ResponseWriter, r *http.Request) {
	claims := jwt_auth.TokenGetClaims(jwtauth.TokenFromHeader(r))

	if claims == nil {
		w.Write([]byte("Invalid token"))
		return
	}

	if claims["role"] != "admin" {
		w.Write([]byte("You don't have permission to access this resource"))
		return
	}

	db := DbConnection()
	var user model.User

	db.Where("username = ?", chi.URLParam(r, "name")).First(&user)

	if user.Username == "" {
		w.Write([]byte("User not found"))
		return
	}

	userResponseDto := model.UserToUserResponseDTO(user)

	data, _ := json.Marshal(userResponseDto)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	claims := jwt_auth.TokenGetClaims(jwtauth.TokenFromHeader(r))

	if claims == nil {
		w.Write([]byte("Invalid token"))
		return
	}

	db := DbConnection()
	var user model.User

	db.Where("username = ?", claims["username"]).First(&user)
	data, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
