package controllers

import (
	"fmt"
	"log"
	"strconv"
	"net/http"
	"encoding/json"
	"strings"

	"store/models"
	"store/db"

	"github.com/gorilla/mux"
)

func ProductosTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Println("El metodo es", r.Method)

	productos := [100]models.Producto{}

	con := db.Conexion{"root", "root", "store"}
	con.AbrirConexion()
	com := con.ObtenerConexion()

	// insert
	query := "SELECT * FROM producto"

	rows, err := com.Query(query)
	if err != nil {
		panic(err)
	}

	var i int = 0;

	for rows.Next(){
		rows.Scan(&productos[i].Id, &productos[i].Nombre, &productos[i].Descripcion, &productos[i].Slug ,&productos[i].Precio, &productos[i].Imagen)
		i++;
	}

	con.CerrarConexion()

	json.NewEncoder(w).Encode(productos)
}

// ProductoController es la funcion que registra nuevos productos
//y devuelve los ya existentes a traves de su id y su slug.
func ProductoController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
		case "POST":
			log.Println("El metodo es", r.Method)

			producto := models.Producto{}

			log.Println(r.Form)

			producto.Nombre = r.PostFormValue("nombre")
			producto.Descripcion = r.PostFormValue("descripcion")
			producto.Slug = strings.Replace(producto.Nombre, " ", "-", 10)
			producto.Precio, _ = strconv.Atoi(r.PostFormValue("precio"))
			producto.Imagen = r.FormValue("imagen")

			con := db.Conexion{"root", "root", "store"}
			con.AbrirConexion()
			com := con.ObtenerConexion()

			// insert
			stmt, err := com.Prepare("INSERT producto SET id=0 ,nombre=?, descripcion=?, slug=?, precio=?, imagen=?")
			if err != nil {
				panic(err)
			}
	
			res, err := stmt.Exec(producto.Nombre, producto.Descripcion, producto.Slug, producto.Precio, producto.Imagen)
			producto.Id, _ = res.LastInsertId()

			if err != nil {
				panic(err)
			}

			con.CerrarConexion()

			json.NewEncoder(w).Encode(producto)
		case "GET":
			log.Println("El metodo es", r.Method)

			vars := mux.Vars(r)

			producto := models.Producto{}

			producto.Id, _ = strconv.ParseInt(vars["id"], 10, 64)
			producto.Slug = vars["slug"]

			con := db.Conexion{"root", "root", "store"}
			con.AbrirConexion()
			com := con.ObtenerConexion()

			// insert
			query := "SELECT * FROM producto WHERE id=" + strconv.Itoa(int(producto.Id)) + " and slug='" + producto.Slug + "'"

			log.Println("EL QUERY ES: ", query)

			rows, err := com.Query(query)
			if err != nil {
				panic(err)
			}
	
			for rows.Next(){
				rows.Scan(&producto.Id, &producto.Nombre, &producto.Descripcion, &producto.Slug ,&producto.Precio, &producto.Imagen)
			}

			con.CerrarConexion()

			json.NewEncoder(w).Encode(producto)
		
	}
}

// ProductoEnvioController es la funcion que registra y envia el envio
//de un determinado producto
func ProductoEnvioController(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL, r.Form)

	fmt.Fprintf(w, "Vista del envio producto")
}

func ProductoIdController(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Println("El metodo es", r.Method)

	producto := models.Producto{}
	
	con := db.Conexion{"root", "root", "store"}
	con.AbrirConexion()
	com := con.ObtenerConexion()

	query1 := "SELECT * FROM producto WHERE id=" + r.PostFormValue("id")

	log.Println(query1)

	rows1, err := com.Query(query1)
	if err != nil {
		panic(err)
	}

	for rows1.Next(){
		rows1.Scan(&producto.Id, &producto.Nombre, &producto.Descripcion, &producto.Slug ,&producto.Precio, &producto.Imagen)
	}
}

func ProductoBorrarId(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Println("El metodo es", r.Method)

	vars := mux.Vars(r)

	producto := models.Producto{}

	producto.Id, _ = strconv.ParseInt(vars["id"], 10, 64)
	producto.Slug = vars["slug"]

	con := db.Conexion{"root", "root", "store"}
	con.AbrirConexion()
	com := con.ObtenerConexion()

	stmt, err := com.Prepare("DELETE FROM pedido WHERE id=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(producto.Id)
	if err != nil {
		panic(err)
	}

	log.Println(res)

	con.CerrarConexion()
}

func ProductoBusqueda(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Println("El metodo es", r.Method)

	productos := [100]models.Producto{}

	termino := mux.Vars(r)["termino"]

	con := db.Conexion{"root", "root", "store"}
	con.AbrirConexion()
	com := con.ObtenerConexion()

	// insert
	query := "SELECT * FROM producto WHERE nombre like '%"+ termino +"%'"

	rows, err := com.Query(query)
	if err != nil {
		panic(err)
	}

	var i int = 0;

	for rows.Next(){
		rows.Scan(&productos[i].Id, &productos[i].Nombre, &productos[i].Descripcion, &productos[i].Slug ,&productos[i].Precio, &productos[i].Imagen)
		i++;
	}

	con.CerrarConexion()

	json.NewEncoder(w).Encode(productos)
}

func ProductoBorrarId1(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Println("El metodo es", r.Method)

	vars := mux.Vars(r)

	producto := models.Producto{}

	producto.Id, _ = strconv.ParseInt(vars["id"], 10, 64)
	producto.Slug = vars["slug"]

	con := db.Conexion{"root", "root", "store"}
	con.AbrirConexion()
	com := con.ObtenerConexion()

	stmt, err := com.Prepare("DELETE FROM producto WHERE id=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(producto.Id)
	if err != nil {
		panic(err)
	}

	log.Println(res)

	con.CerrarConexion()
}

func ProductoActualizar(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println("El metodo es", r.Method)

	producto := models.Producto{}

	log.Println(r.Form)

	producto.Id,_ = strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	producto.Nombre = r.PostFormValue("nombre")
	producto.Descripcion = r.PostFormValue("descripcion")
	producto.Slug = strings.Replace(producto.Nombre, " ", "-", 10)
	producto.Precio, _ = strconv.Atoi(r.PostFormValue("precio"))
	producto.Imagen = r.FormValue("imagen")

	log.Println(producto)

	con := db.Conexion{"root", "root", "store"}
	con.AbrirConexion()
	com := con.ObtenerConexion()

	// insert
	stmt, err := com.Prepare("UPDATE producto SET nombre=?, descripcion=?, slug=?, precio=?, imagen=? WHERE id=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(producto.Nombre, producto.Descripcion, producto.Slug, producto.Precio, producto.Imagen, producto.Id)

	if err != nil {
		panic(err)
	}

	log.Println(res)

	con.CerrarConexion()

	json.NewEncoder(w).Encode(producto)
}