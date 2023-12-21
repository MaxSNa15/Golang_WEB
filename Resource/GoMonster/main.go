package main

import (
	"GoMonster/handlers"
	"log"
	"net/http"
)


func main() {
	// Crear un Enrutador
	router := http.NewServeMux()

	// Configurar los archivos estáticos
	fileSta := http.FileServer(http.Dir("static"))
	// Configurar la ruta para los archivos estáticos
	router.Handle("/static/", http.StripPrefix("/static/", fileSta))

	// Configurar las rutas
	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/new", handlers.NewGameHandler)
	router.HandleFunc("/game", handlers.GameHandler)
	router.HandleFunc("/play", handlers.PlayHandler)
	router.HandleFunc("/about", handlers.AboutHandler)

	// Crear el servidor
	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}