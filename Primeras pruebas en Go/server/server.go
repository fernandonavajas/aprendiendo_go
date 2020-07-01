package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
		//http.ServeFile(w, r, "index.html") otra forma de subir solo 1 archivo
	})
	http.ListenAndServe(":8000", nil)
}
