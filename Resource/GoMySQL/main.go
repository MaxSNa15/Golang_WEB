package main

import (
	"GoMySQL/data"
	"GoMySQL/models"
	"fmt"
)

func main() {
	data.Connect()
	// data.CreateTable(models.UserSchema, "users")
	// data.TruncateTable("users") // Elimina todos los registros de la tabla
	data.Ping()
	// userJua := models.CreateUser("Roel", "1234567", "roel@ca.com")
	// fmt.Println(userJua)

	userss := models.ListUsers()
	fmt.Println(userss)

	user := models.GetUser(2)
	fmt.Println(user)
	// Actualizar un usuario
	user.Name = "SNava"
	user.Save()
	fmt.Println(user)

	// Eliminar un usuario
	user.Delete()

	fmt.Println(models.ListUsers())

	data.TruncateTable("users") // Elimina todos los registros de la tabla

	fmt.Println(models.ListUsers())


	data.Close()
	// data.Ping() // panic: sql: database is closed
}

