package main

import (
	"go-api/adaptors/db"
	"go-api/adaptors/rest"
	"go-api/domain/service"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		rest.WriteJSONResponse(w, http.StatusOK, &map[string]string{"status": "ok", "version": "0.0.1"})
	})

	dbAdapter := db.NewBun()

	userRepo := db.NewUserRepository(dbAdapter)

	userService := service.NewUserService(userRepo)
	UserHandler := rest.NewUserHandler(userService)

	rest.RegisterCRUD(mux, "/users", UserHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", rest.CORSHandler(mux)))
}
