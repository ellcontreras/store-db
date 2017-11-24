package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL, r.Form)

	fmt.Fprintf(w, "Todos los productos")
}
