package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// Creamos una estructura para la pagina
type Page struct {
	Title string
	Body []byte
}

// Creamos un metodo para guardar la pagina
func (p *Page) save() error { 
	filename := p.Title + ".txt" // Creamos el nombre del archivo con extencion y lo guardamos en filename
	return os.WriteFile(filename, p.Body, 0600) // Guardamos el archivo con los permisos 0600 (solo lectura)
}

// Creamos un metodo para cargar la pagina
func loadPage(title string) (*Page, error) {
	filename := title + ".txt" // Creamos el nombre del archivo con extencion y lo guardamos en filename
	// Creamos dos variables, body y err, que recibiran el contenido del archivo y el error si existe
	body, err := os.ReadFile(filename)
	if err != nil { // Si existe un error
		return nil, err // Retornamos nil y el error
	}
	return &Page{Title: title, Body: body}, nil // Retornamos la pagina y nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html") // Cargamos el archivo edit.html
	t.Execute(w, p) // Ejecutamos el archivo edit.html
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	// Cargar la pagina
	title := r.URL.Path[len("/view/"):] // Obtenemos el titulo de la pagina
	// Cargar la pagina
	p, err := loadPage(title) // Cargamos la pagina
	if err != nil { // Si existe un error
		http.Redirect(w, r, "/edit/"+title, http.StatusFound) // Redireccionamos a la pagina de edicion
		return // Retornamos
	}
	//fmt.Fprintf(w, "<h1>%s</h1> <div>%s</div>", p.Title, p.Body) // Imprimimos el mensaje
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/edit/"):] // Obtenemos el titulo de la pagina
	p, err := loadPage(title) // Cargamos la pagina
	if err != nil { // Si existe un error
		p = &Page{Title: title} // Creamos una nueva pagina
	}
	renderTemplate(w, "edit", p)
}

func main() {
	// Crear rutas 
	// Responder al clinete con un mensaje
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)

	// Levantar el servidor
	log.Fatal(http.ListenAndServe(":8080", nil)) // Escuchamos en el puerto 8080
}