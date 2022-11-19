package main

import (
	//Para la pagina wed
	"database/sql"
	"log"
	"net/http"
	//buscar registros o plantillas atravez de carpetas
	"text/template"
	//drivers
	//funcion para conexion de la base de datos

	_ "github.com/go-sql-driver/mysql"

	"sistema/contacto"

)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func conexionDB() *sql.DB {
	conexion, err := sql.Open("mysql", "root:"+""+"@tcp(localhost:3306)/agenda")
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

func main() {
	http.HandleFunc("/", Inicio)
	log.Print("servidor funcionando")
	http.ListenAndServe(":5500", nil)
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	conexionEstablecida := conexionDB()
	registros, err := conexionEstablecida.Query("SELECT * FROM recordarcontacto")
	if err != nil {
		panic(err.Error())
	}
	empleado := contacto.Contacto{}
	AregloEmpleado := []contacto.Contacto{}
	for registros.Next() {
		var id int
		var nombre, actividad, lugar, dia, hora string
		err = registros.Scan(&id, &nombre, &actividad, &lugar, &hora, &dia)
		//enmarcar los errores
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Actividad = actividad
		empleado.Lugar = lugar
		empleado.Hora = hora
		empleado.Dia = dia
		AregloEmpleado = append(AregloEmpleado, empleado)
	}
	plantillas.ExecuteTemplate(w, "Inicio", AregloEmpleado)

}
