package controllers

import (
	"log"
	"net/http"

	"store/models"
	"store/db"

	"encoding/json"
	"github.com/gorilla/mux"

	"strconv"
)

type OrdenPedidos struct {
	Id int64
	Id_usuario int64
	NombreProducto string
	Fecha string
	IdP string
}

func PedidoController(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Println("El metodo es", r.Method)

	pedido := models.Pedido{}

	log.Println(r.Form)

	idUsuario, _ := mux.Vars(r)["idUsuario"];
	idproducto,_ := mux.Vars(r)["idProducto"]

	i, _ := strconv.Atoi(idUsuario)
	i1, _ := strconv.Atoi(idproducto)

	pedido.Id_usuario = int64(i)
	pedido.Id_producto = int64(i1)

	con := db.Conexion{"root", "root", "store"}
	con.AbrirConexion()
	com := con.ObtenerConexion()

	// insert
	stmt, err := com.Prepare("INSERT pedido SET id_usuario=?, id_producto=?, fecha=Now()")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(pedido.Id_usuario, pedido.Id_producto)
	pedido.Id, _ = res.LastInsertId()

	if err != nil {
		panic(err)
	}

	con.CerrarConexion()

	json.NewEncoder(w).Encode(pedido)
}

func PedidosTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method{
		case "GET":
			w.Header().Set("Access-Control-Allow-Origin", "*")

			log.Println("El metodo es", r.Method)

			pedidos := [100]OrdenPedidos{}

			con := db.Conexion{"root", "root", "store"}
			con.AbrirConexion()
			com := con.ObtenerConexion()

			// insert
			query := "SELECT pedido.id, pedido.id_usuario,pedido.fecha, producto.nombre, producto.id FROM pedido right JOIN producto ON pedido.id_producto = producto.id where pedido.id_usuario =" + mux.Vars(r)["idUsuario"]

			rows, err := com.Query(query)
			if err != nil {
				panic(err)
			}

			var i int = 0;

			for rows.Next(){
				rows.Scan(&pedidos[i].Id, &pedidos[i].Id_usuario, &pedidos[i].NombreProducto, &pedidos[i].Fecha, &pedidos[i].IdP)
				i++;
			}
			con.CerrarConexion()

			json.NewEncoder(w).Encode(pedidos)
	}
}