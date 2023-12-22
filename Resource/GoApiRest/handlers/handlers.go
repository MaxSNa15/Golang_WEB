package handlers

import (
	"GoApiRest/data"
	"GoApiRest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Funcion para obtener todos los usuarios
func GetUsers(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(rw, "not implemented yet !")

	rw.Header().Set("Content-Type", "application/json")

	data.Connect()
	users := models.ListUsers()
	data.Close()
	output, _ := json.Marshal(users)
	fmt.Fprintln(rw, string(output))

}

// Funcion para obtener un usuario por su id
func GetUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtener el ID
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	data.Connect()
	user := models.GetUser(userId)
	data.Close()
	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

// Funcion para crear un usuario
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtener todo el registro
	user := models.User{}
	// decodificar de json a struct
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity, err)
	}else{
		data.Connect()
		user.Save()
		data.Close()
	}

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

// Funcion para actualizar un usuario
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtener todo el registro
	user := models.User{}
	// decodificar de json a struct
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity, err)
	}else{
		data.Connect()
		user.Save()
		data.Close()
	}

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

// Funcion para eliminar un usuario
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtener el ID
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	data.Connect()
	user := models.GetUser(userId)
	user.Delete()
	data.Close()
	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}