package routes

import (
	"store/controllers"

	"github.com/gorilla/mux"
)

// Router es la variable es la encargada de manejar las rutas de la app.
var Router *mux.Router

// Init es la funcion que se encarga de iniciar cada ruta de nuestra app.
func Init() {
	Router = mux.NewRouter()

	Router.HandleFunc("/", controllers.HomeController)
	Router.HandleFunc("/usuario/", controllers.UsuarioController).Methods("POST")
	Router.HandleFunc("/usuario/{email}/{contrasena}", controllers.UsuarioController).Methods("GET")

	Router.HandleFunc("/productobusqueda/{termino}", controllers.ProductoBusqueda).Methods("GET")

	Router.HandleFunc("/pedidos/{idUsuario}", controllers.PedidosTodos).Methods("GET")
	Router.HandleFunc("/pedidos/{idUsuario}/{idProducto}", controllers.PedidoController).Methods("POST")
	Router.HandleFunc("/productoborrar/{id}", controllers.ProductoBorrarId)
	Router.HandleFunc("/productoborrar1/{id}", controllers.ProductoBorrarId1)
	Router.HandleFunc("/producto/{id}", controllers.ProductoIdController)
	Router.HandleFunc("/producto", controllers.ProductoController).Methods("POST")
	Router.HandleFunc("/productoActualizar/", controllers.ProductoActualizar).Methods("POST")

	Router.HandleFunc("/productos/", controllers.ProductosTodos)
	Router.HandleFunc("/producto/{id:[0-9]+}/{slug}", controllers.ProductoController).Methods("GET")
}