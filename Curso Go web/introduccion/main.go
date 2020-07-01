package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Nombre", "Valor del header")
		fmt.Fprintf(w, "Hola Mundo por metodo: "+r.Method)
		url := createURL()
		fmt.Println("La url final es: " + url)
		switch r.Method {
		case "GET":
			//r.URL.RawQuery()
			fmt.Println(r.URL.Query())        //params?name=Eduardo ->  devuelve -> map[name:[Eduardo]]
			name := r.URL.Query().Get("name") // -> Eduardo
			if len(name) != 0 {
				fmt.Println(name)
			}
			fmt.Fprintf(w, "metodo: "+r.Method)
		case "POST":
			fmt.Println(r.Header)
			accessToken := r.Header.Get("access_token")
			if len(accessToken) != 0 {
				fmt.Println(accessToken)
			}
			fmt.Fprintf(w, "metodo: "+r.Method)
		case "PUT":
			fmt.Fprintf(w, "metodo: "+r.Method)
		case "DELETE":
			fmt.Fprintf(w, "metodo: "+r.Method)
		default:
			http.Error(w, "Metodo no valido.", http.StatusBadRequest)
		}

	})
	http.HandleFunc("/dos", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/error", http.StatusMovedPermanently)

	})
	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)

	})
	http.HandleFunc("/otro_error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Esto es otro error", http.StatusMovedPermanently)

	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func createURL() string {
	u, err := url.Parse("/params")
	if err != nil {
		panic(err)
	}
	u.Host = "localhost:8000"
	u.Scheme = "http"
	query := u.Query()
	query.Add("nombre", "valor")

	u.RawQuery = query.Encode()

	return u.String()
}
