package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
)

/*
var movies = Movies{
	Movie{"Sin Limites", 2013, "Desconocido"},
	Movie{"Batman Begins", 1999, "Scorsese"},
	Movie{"A todo gas", 2005, "Pizzi"},
}
*/

// Crear variable global
var collection = getSession().DB("curso_go").C("movies")

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
	var results []Movie
	err := collection.Find(nil).Sort("-_id").All(&results)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados: ", results)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	fmt.Fprintf(w, "Has cargado la pelicula numero %s", movie_id)
}

// Recibir un reponseWriter y una request
func MovieAdd(w http.ResponseWriter, r *http.Request) {
	// Recogemos los datos del Body que nos llega en la request
	// Lo decodificamos con JSON
	decoder := json.NewDecoder(r.Body)

	// Creamos una variable para los datos que recogemos
	var movie_data Movie
	// &: Indicar que es una variable que no tiene nada
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}

	// Cerrar la lectura de algo (body)
	defer r.Body.Close()

	//log.Println(movie_data)
	//movies = append(movies, movie_data)

	// Guardar los datos en la base de datos
	// movie_data es el objeto que estamos recibiendo por el body
	err = collection.Insert(movie_data)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(movie_data)
}
