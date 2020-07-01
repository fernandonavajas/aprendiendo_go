package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
	vars := mux.Vars(r)
	nombre := vars["nombre"]

	fmt.Fprintf(w, "Los parametros son: "+nombre)
}

func twoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Estoy en el 2do handler")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/usuarios/{nombre}", YourHandler) //curl -i http://localhost:8000/usuarios/ricardo
	r.HandleFunc("/2", twoHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
