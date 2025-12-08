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
	rest.RegisterCRUD("/users", UserHandler)

	http.ListenAndServe(":8080", nil)
}
