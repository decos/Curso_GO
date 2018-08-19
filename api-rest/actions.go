package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
)

var movies = Movies{
	Movie{"Sin Limites", 2013, "Desconocido"},
	Movie{"Batman Begins", 1999, "Scorsese"},
	Movie{"A todo gas", 2005, "Pizzi"},
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return session
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*
func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esta es la pagina de contacto")
}
*/

func MovieList(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(movies)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	fmt.Fprintf(w, "Has cargado la pelicula numero %s", movie_id)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	// &: Indicar que es una variable que no tiene nada
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}

	// Cerrar la lectura de algo (body)
	defer r.Body.Close()

	log.Println(movie_data)
	movies = append(movies, movie_data)
}
