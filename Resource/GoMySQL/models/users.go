package models

import "GoMySQL/data"

type User struct {
	Id 	 		int64
	Name 		string
	password 	string
	email 		string
}

type Users []User

// Query para insertar un usuario
const UserSchema string = 
`CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(30) NOT NULL,
	password VARCHAR(30) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP)`

// Constructor de la estructura User
func NewUser(name, password, email string) *User {
	user := &User{
		Name: name,
		password: password,
		email: email,
	}
	return user
}

// Insertar un Registro en la base de datos
func (this *User) insert(){
	sql := "INSERT users SET name=?, password=?, email=?"
	result, _ := data.Exec(sql, this.Name, this.password, this.email)
	// Obtenemos el ultimo id insertado
	this.Id, _ = result.LastInsertId()
}

// Funcion para crear un usuario en la base de datos
func CreateUser(name, password, email string) *User {
	user := NewUser(name, password, email)
	user.Save()
	return user
}

// Listar todos los usuarios de la base de datos
func ListUsers() Users {
	sql := "SELECT id, name, password, email FROM users"
	users := Users{}
	rows, _ := data.Query(sql)
	for rows.Next() {
		// Creamos un nuevo usuario (objeto)
		user := User{}
		// Escaneamos los datos de la fila y los guardamos en el objeto
		rows.Scan(&user.Id, &user.Name, &user.password, &user.email)
		// Agregamos el usuario a la lista de usuarios
		users = append(users, user)
	}
	return users
}

// Obtener un registro de la base de datos
func GetUser(id int) *User {
	user := NewUser("", "", "")

	sql := "SELECT id, name, password, email FROM users WHERE id = ?"
	rows, _ := data.Query(sql, id)
	for rows.Next() {
		// Escaneamos los datos de la fila 
		rows.Scan(&user.Id, &user.Name, &user.password, &user.email)
	}
	return user
}

// Actualizar un registro de la base de datos
func (this *User) update() {
	sql := "UPDATE users SET name=?, password=?, email=? WHERE id=?"
	_, _ = data.Exec(sql, this.Name, this.password, this.email, this.Id)
}

/*
	! Funcion GLOBAL que hace el INSERT y UPDATE
	verifica si el id ya existe en la base de datos hace un UPDATE
	Si no existe hace un INSERT
*/
func (this *User) Save() {
	if this.Id == 0 {
		this.insert()
	} else {
		this.update()
	}
}

// Eliminar un registro de la base de datos
func (this *User) Delete() {
	sql := "DELETE FROM users WHERE id=?"
	_, _ = data.Exec(sql, this.Id)
}