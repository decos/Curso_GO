package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*
func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esta es la pagina de contacto")
}
*/

func MovieList(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{"Sin Limites", 2013, "Desconocido"},
		Movie{"Batman Begins", 1999, "Scorsese"},
		Movie{"A todo gas", 2005, "Pizzi"},
	}

	json.NewEncoder(w).Encode(movies)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	fmt.Fprintf(w, "Has cargado la pelicula numero %s", movie_id)
}
