package main

import (
	"encoding/json"
	"net/http"
)

type Curso struct {
	Title        string `json:"titulo"`
	NumeroVideos int    `json:"numero_videos"`
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		cursos := Cursos{
			Curso{"Curso de Go", 20},
			Curso{"Curso de Python", 30},
			Curso{"Curso de Java", 40},
		}
		json.NewEncoder(w).Encode(cursos)
	})
	http.ListenAndServe(":8001", nil)
}
