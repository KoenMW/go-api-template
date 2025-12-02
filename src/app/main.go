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
	http.HandleFunc("/users", UserHandler.CreateUser)
	http.HandleFunc("/users/", UserHandler.GetUser)

	http.ListenAndServe(":8080", nil)
}
