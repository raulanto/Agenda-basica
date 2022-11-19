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
	
)
var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func conexionDB() *sql.DB {
	conexion, err := sql.Open("mysql", "root:"+""+"@tcp(localhost:3306)/sistema")
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
	plantillas.ExecuteTemplate(w, "Inicio", nil)
}