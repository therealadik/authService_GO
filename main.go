package main

import (
	"log"
	"net/http"

	"github.com/therealadik/auth-service/internal/database"
	"github.com/therealadik/auth-service/internal/handlers"
)

func main() {
	err := database.InitDB()
	if err != nil{
		log.Fatal(err.Error())
		return
	}

	http.HandleFunc("/auth", handlers.AuthHandler)
	http.HandleFunc("/refresh", handlers.RefreshTokens)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
