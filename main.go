package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"

)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func conexionDB() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contraseña := ""
	Nombre := "sistema"

	conexion, err := sql.Open(Driver, Usuario+":"+Contraseña+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion

}

type Empleado struct {
	Id        int
	Nombre    string
	Email     string
	Documento int
	Telefono  int
}

func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/insertar", Insertar)
	http.HandleFunc("/borrar", Borrar)
	http.HandleFunc("/editar", Editar)
	http.HandleFunc("/actualizar", Actualizar)

	log.Println("Servidor corrriendo")
	http.ListenAndServe(":8080", nil)
}

func Borrar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	fmt.Println(idEmpleado)
	conexionEstablecida := conexionDB()
	borrarRegistro, err := conexionEstablecida.Prepare("Delete from empleados where id=?")
	if err != nil {
		panic(err.Error())
	}
	borrarRegistro.Exec(idEmpleado)
	http.Redirect(w, r, "/", 301)

}

func Inicio(w http.ResponseWriter, r *http.Request) {

	//conexion a base de datos
	conexionEstablecida := conexionDB()
	registros, err := conexionEstablecida.Query("SELECT * from empleados")
	if err != nil {
		panic(err.Error())
	}
	empleado := Empleado{}
	arrayEmpleado := []Empleado{}

	for registros.Next() {
		var id, documento, telefono int
		var nombre, email string
		err = registros.Scan(&id, &nombre, &email, &documento, &telefono)
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Email = email
		empleado.Documento = documento
		empleado.Telefono = telefono

		arrayEmpleado = append(arrayEmpleado, empleado)

	}

	plantillas.ExecuteTemplate(w, "inicio", arrayEmpleado)

}
func Editar(w http.ResponseWriter, r *http.Request) {

	idEmpleado := r.URL.Query().Get("id")
	conexionEstablecida := conexionDB()
	registro, err := conexionEstablecida.Query("SELECT * FROM empleados where id=?", idEmpleado)
	empleado := Empleado{}

	for registro.Next() {
		var id,documento,telefono int
		var nombre, email string
		err = registro.Scan(&id, &nombre, &email,&documento,&telefono)
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Email = email
		empleado.Documento = documento
		empleado.Telefono = telefono

	}
	fmt.Println(empleado)
	plantillas.ExecuteTemplate(w, "editar", empleado)

}

func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)

}

func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		email := r.FormValue("email")
		documento := r.FormValue("documento")
		telefono := r.FormValue("telefono")

		//conexion a base de datos
		conexionEstablecida := conexionDB()
		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleados(nombre,email,documento,telefono) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(nombre, email, documento, telefono)
		http.Redirect(w, r, "/", 301)

	}
}
func Actualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nombre := r.FormValue("nombre")
		email := r.FormValue("email")
		documento := r.FormValue("documento")
		telefono := r.FormValue("telefono")

		//conexion a base de datos
		conexionEstablecida := conexionDB()
		actualizarRegistros, err := conexionEstablecida.Prepare("UPDATE empleados SET nombre=?,email=?,documento=?,telefono=? WHERE id =?")
		if err != nil {
			panic(err.Error())
		}
		actualizarRegistros.Exec(nombre, email, documento, telefono, id)
		http.Redirect(w, r, "/", 301)

	}
}
