package main

import (
	"github.com/estudodb/entities"
	"github.com/gorilla/mux"

)

var movies = []entities.Movie
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}",getMovie).Methods("GET")
}
