package main

import (
	"fmt"
	"net/http"
	"store/config"
	"store/routes"
)

func main() {
	//Inicializamos las rutas
	routes.Init()

	//Inicializamos la configuracion
	config.Init()

	//Escuchar servidor
	fmt.Println("El servidor está corriendo en localhost" + config.Config["puerto"])
	http.ListenAndServe(config.Config["puerto"], routes.Router)
}
