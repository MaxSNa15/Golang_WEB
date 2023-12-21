package handlers

import (
    "fmt"
    "net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Pagina de Inicio de GoMonster")
}

func NewGameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pagina de Nuevo Juego de GoMonster")
}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pagina de Juego de GoMonster")
}

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Juego de GoMonster")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pagina de Acerca de GoMonster")
}