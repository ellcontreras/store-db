package controllers

import (
	"log"
	"net/http"

	"store/models"
	"store/db"

	"encoding/json"
	"github.com/gorilla/mux"
)

func UsuarioController(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
		case "POST":
			log.Println("El metodo es", r.Method)

			usuario := models.Usuario{}

			log.Println(r.Form)

			usuario.Nombre = r.PostFormValue("nombre")
			usuario.Apellidos = r.PostFormValue("apellidos")
			usuario.Email = r.PostFormValue("email")
			usuario.Contrasena = r.PostFormValue("contrasena")

			con := db.Conexion{"root", "root", "store"}
			con.AbrirConexion()
			com := con.ObtenerConexion()

			// insert
			stmt, err := com.Prepare("INSERT usuario SET id=0 ,nombre=?, apellidos=?, email=?, contrasena=?")
			if err != nil {
				panic(err)
			}
	
			res, err := stmt.Exec(usuario.Nombre, usuario.Apellidos, usuario.Email, usuario.Contrasena)
			usuario.Id, _ = res.LastInsertId()

			if err != nil {
				panic(err)
			}

			con.CerrarConexion()

			json.NewEncoder(w).Encode(usuario)
		case "GET":
			log.Println("El metodo es", r.Method)
			vars := mux.Vars(r)

			usuario := models.Usuario{}

			usuario.Email, _ = vars["email"]
			usuario.Contrasena = vars["contrasena"]

			con := db.Conexion{"root", "root", "store"}
			con.AbrirConexion()
			com := con.ObtenerConexion()

			// insert
			query := "SELECT * FROM usuario WHERE email='" + usuario.Email + "' and contrasena='" + usuario.Contrasena + "'"

			log.Println("EL QUERY ES: ", query)

			rows, err := com.Query(query)
			if err != nil {
				panic(err)
			}
	
			for rows.Next(){
				rows.Scan(&usuario.Id, &usuario.Nombre, &usuario.Apellidos, &usuario.Email ,&usuario.Contrasena)
			}

			con.CerrarConexion()

			json.NewEncoder(w).Encode(usuario)
	}
}