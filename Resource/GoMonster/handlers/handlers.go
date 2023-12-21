package handlers

import (
    "fmt"
	"html/template"
    "net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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