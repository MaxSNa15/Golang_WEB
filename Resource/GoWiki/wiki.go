package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
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

var templates = template.Must(template.ParseFiles("edit.html", "view.html")) // Cargamos los archivos edit.html y view.html

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p) // Ejecutamos el template
	if err != nil { // Si existe un error
		http.Error(w, err.Error(), http.StatusInternalServerError) // Creamos un error interno
		return // Retornamos
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$") // Creamos una expresion regular para validar la ruta


func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path) // Buscamos la ruta en la expresion regular
		if m == nil { // Si no existe
			http.NotFound(w, r) // Creamos un error 404
			return
		}
		fn(w, r, m[2]) // Ejecutamos la funcion
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string){
	p, err := loadPage(title) // Cargamos la pagina
	if err != nil { // Si existe un error
		http.Redirect(w, r, "/edit/"+title, http.StatusFound) // Redireccionamos a la pagina de edicion
		return // Retornamos
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string){
	p, err := loadPage(title) // Cargamos la pagina
	if err != nil { // Si existe un error
		p = &Page{Title: title} // Creamos una nueva pagina
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string){
	body := r.FormValue("body") // Obtenemos el contenido de la pagina
	p := &Page{Title: title, Body: []byte(body)} // Creamos una nueva pagina
	err := p.save() // Guardamos la pagina
	if err != nil { // Si existe un error
		http.Error(w, err.Error(), http.StatusInternalServerError) // Creamos un error interno
		return // Retornamos
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound) // Redireccionamos a la pagina de edicion
}

func main() {
	// Crear rutas 
	// Responder al clinete con un mensaje
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	// Levantar el servidor
	log.Fatal(http.ListenAndServe(":8080", nil)) // Escuchamos en el puerto 8080
}