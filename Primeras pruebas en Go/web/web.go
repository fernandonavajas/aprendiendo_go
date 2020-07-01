package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Soy el usuario")
	})
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hay una nueva peticion")
	io.WriteString(w, "Que estas buscando")
}
