package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ProductoController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	log.Println(r.Method, r.URL, r.Form)

	fmt.Fprintf(w, "Vista de producto "+vars["id"])
}

func ProductoEnvioController(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL, r.Form)

	fmt.Fprintf(w, "Vista del envio producto")
}
