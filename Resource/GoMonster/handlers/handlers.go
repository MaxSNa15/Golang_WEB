package handlers

import (
	"GoMonster/logic"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

const (
	tmplDir = "templates/"
	baseTmpl = tmplDir + "base.html"
)

type Player struct {
	Name string
}

var player Player

func renderTemplate(w http.ResponseWriter, base, page string, data any) {
	tpl := template.Must(template.ParseFiles(base, tmplDir+page+".html"))
	
	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	Reset()
	renderTemplate(w, baseTmpl, "index", nil)
}

func NewGameHandler(w http.ResponseWriter, r *http.Request) {
	Reset()
	renderTemplate(w, baseTmpl, "new-game", nil)
}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		player.Name = r.Form.Get("name")
	}
	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
	}
	renderTemplate(w, baseTmpl, "game", player)
}

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := logic.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	Reset()
	renderTemplate(w, baseTmpl, "about", nil)
}

// Reiniciar valores
func Reset() {
	player.Name = ""
	logic.ComputerScore = 0
	logic.PlayerScore = 0
}