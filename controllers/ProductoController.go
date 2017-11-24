package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func ProductoController(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL, r.Form)

	fmt.Fprintf(w, "Vista de producto")
}

func ProductoEnvioController(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL, r.Form)

	fmt.Fprintf(w, "Vista del envio producto")
}
