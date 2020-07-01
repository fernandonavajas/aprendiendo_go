package main

import (
	"fmt"
	"log"
	"net/http"

	"./mux"
)

func Hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola desde una funcion")
}

type User struct {
	name string
}

func (this *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola desde un handler "+this.name)
}

func main() {
	user := &User{"Eduardo"}
	mux := mux.CreateMux()

	mux.AddFunc("/hola", Hola)
	mux.AddHandle("/usuario", user)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
