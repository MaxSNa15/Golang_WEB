package main

import (
	"GORM/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// models.MigrarUser()

	// Rutas
	mux := mux.NewRouter()

	// Endpoints
	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", mux))
}