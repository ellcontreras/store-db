package controllers

import (
	"fmt"
	"log"
	"net/http"
)

// HomeController es la funcion que se encarga de mandar la info a la pagina principal
func HomeController(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL, r.Form)

	fmt.Fprintf(w, "Todos los productos")
}
