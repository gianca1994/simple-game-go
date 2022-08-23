package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"net/http"
	"simple-game-golang/src/cmd/routes"
)

func main() {

	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Get("/", routes.GetAllUsers)
		r.Get("/{name}", routes.GetUserByName)
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", routes.Login)
		r.Post("/register", routes.Register)
	})

	http.ListenAndServe(":8080", r)

	// var user1, user2 model.User
	// db.First(&user1, "username = ?", "gianca")
	// db.First(&user2, "username = ?", "lucho")

	// service.CombatSystem(user1, user2)
}
