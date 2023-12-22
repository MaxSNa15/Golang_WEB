package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// username:password@tcp(host:port)/database-name
const url = "root:12345@tcp(localhost:3306)/goweb"

// variable global
var data *sql.DB

// Funcion conecta a la base de datos
func Connect() {
	conn, err := sql.Open("mysql", url)
	if err != nil {
		panic(err.Error())
	}
	data = conn
	fmt.Println("Conectado a la base de datos")
}

// Funcion que cierra la conexion a la base de datos
func Close() {
	data.Close()
	fmt.Println("Desconectado de la base de datos")
}
