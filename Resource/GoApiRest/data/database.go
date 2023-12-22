package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// username:password@tcp(host:port)/database-name
const url = "root:12345@tcp(localhost:3306)/goweb"

// variable global que guarda la concexion a la base de datos
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

// Verifica la conexion a la base de datos
func Ping() {
	if err := data.Ping(); err != nil {
		panic(err.Error())
	}
	fmt.Println("Ping a la base de datos")
}

// Verificar si una tabla existe en la base de datos
func TableExists(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(sql)
	if err != nil {
		panic(err.Error())
	}
	return rows.Next()
}

// Crear una tabla en la base de datos
func CreateTable(schema string, nameTable string) {
	if !TableExists(nameTable) {
		_ , err := Exec(schema)
		if err != nil {
			panic(err.Error())
		}
	}
}

// Funcion para truncar una tabla (eliminar todos los registros de una tabla)
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE TABLE %s", tableName)
	_ , err := Exec(sql)
	if err != nil {
		panic(err.Error())
	}
}

/*
	Estas funciones son polimorficas, es decir, pueden recibir cualquier tipo de dato
	se usan para ejecutar directamente las funciones sin necesidad de un modulo 'data.Exec()'
	directamente Exec() o Query()
	Tambien se pueden usar en otros paquetes
*/

// Polimorfismo de la funcion Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := data.Exec(query, args...)
	Close()
	if err != nil {
		panic(err.Error())
	}
	return result, err
}

// Polimorfismo de la funcion Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := data.Query(query, args...)
	Close()
	if err != nil {
		panic(err.Error())
	}
	return rows, err
}
