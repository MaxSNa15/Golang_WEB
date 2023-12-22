package data

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Realiza la conexion
var dsn = "root:12345@tcp(localhost:3306)/goweb?charset=utf8mb4&parseTime=True&loc=Local"
var Database = func() (data *gorm.DB) {
	if data, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la conexion", err)
		panic(err)
	} else {
		fmt.Println("Conexion exitosa")
		return data
	}
}()
