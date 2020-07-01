package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}*/

type User struct {
	name string
}

func HolaDos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HolaDos")
}

func HolaUno(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HolaUno")
}

func (this *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola: "+this.name)
}

func main() {
	eduardo := &User{name: "Eduardo"}
	redirect := http.RedirectHandler("https://google.com", http.StatusMovedPermanently)
	notFound := http.NotFoundHandler()
	mux := http.NewServeMux()
	mux.HandleFunc("/1", HolaUno)
	mux.HandleFunc("/2", HolaDos)
	mux.HandleFunc("/3", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola 3 (anonima)")
	})
	mux.Handle("/", eduardo)
	mux.Handle("/moved", redirect)
	mux.Handle("/notFound", notFound)
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
