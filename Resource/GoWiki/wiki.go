package main

import (
	"fmt"
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

func viewHandler(w http.ResponseWriter, r *http.Request){
	// Cargar la pagina
	title := r.URL.Path[len("/view/"):] // Obtenemos el titulo de la pagina
	// Cargar la pagina
	p, _ := loadPage(title) // Cargamos la pagina
	fmt.Fprintf(w, "<h1>%s</h1> <div>%s</div>", p.Title, p.Body) // Imprimimos el mensaje
}

func main() {
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")} // Creamos una pagina
	// p1.save() // Guardamos la pagina
	// p2, _ := loadPage("TestPage") // Cargamos la pagina
	// fmt.Println(string(p2.Body)) // Imprimimos el contenido de la pagina

	// Responder al clinete con un mensaje
	http.HandleFunc("/view/", viewHandler)

	// Levantar el servidor
	log.Fatal(http.ListenAndServe(":8080", nil)) // Escuchamos en el puerto 8080
}