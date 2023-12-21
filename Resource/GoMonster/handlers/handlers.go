package handlers

import (
    "fmt"
	"html/template"
    "net/http"
)

const (
	tmplDir = "templates/"
	baseTmpl = tmplDir + "base.html"

)

func renderTemplate(w http.ResponseWriter, base, page string, data any) {
	tpl := template.Must(template.ParseFiles(base, tmplDir+page+".html"))
	
	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, baseTmpl, "index", nil)
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