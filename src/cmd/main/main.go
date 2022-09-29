package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	_ "github.com/lib/pq"
	"net/http"
	"simple-game-golang/src/cmd/routes"
)

var tokenAuth *jwtauth.JWTAuth

func main() {
	port := ":8080"
	fmt.Printf("Starting server on %v\n", port)

	// database.NewPostgreSQL() => Migration table users

	http.ListenAndServe(port, routerHandler())
}

func routerHandler() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!!"))
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", routes.GetAllUsers)
		r.Get("/{name}", routes.GetUserByName)
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", routes.Login)
		r.Post("/register", routes.Register)
	})

	r.Group(func(r chi.Router) {
		r.Get("/profile", routes.GetProfile)
	})

	return r
}
