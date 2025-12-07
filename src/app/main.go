package main

import (
	"go-api/adaptors/db"
	"go-api/adaptors/rest"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var bun, err = db.NewBun()
	if err != nil {
		panic(err)
	}

	UserHandler := rest.NewUserHandler(bun)
	http.HandleFunc("POST /users", UserHandler.CreateUser)
	http.HandleFunc("GET /users/{id}", UserHandler.GetUser)
	http.HandleFunc("GET /users", UserHandler.ListUsers)
	http.HandleFunc("PUT /users/{id}", UserHandler.UpdateUser)
	http.HandleFunc("DELETE /users/{id}", UserHandler.DeleteUser)

	http.ListenAndServe(":8080", nil)
}
