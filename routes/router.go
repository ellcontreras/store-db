package routes

import (
	"store/controllers"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func Init() {
	Router = mux.NewRouter()

	Router.HandleFunc("/", controllers.HomeController)
	Router.HandleFunc("/producto/{id:[0-9]+}/{slug}", controllers.ProductoController)
	Router.HandleFunc("/producto/envio/", controllers.ProductoEnvioController)
}
