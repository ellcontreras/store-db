package main

import (
	"fmt"
	"net/http"
	"store/controllers"
)

func main() {
	//Controladores
	http.HandleFunc("/", controllers.HomeController)
	http.HandleFunc("/producto/", controllers.ProductoController)
	http.HandleFunc("/producto/envio/", controllers.ProductoEnvioController)

	//Escuchar servidor
	fmt.Println("El servidor está corriendo en localhost:8080")
	http.ListenAndServe(":8080", nil)
}
