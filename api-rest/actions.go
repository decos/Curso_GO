package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

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

func responseMovie(w http.ResponseWriter, status int, results Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func responseMovies(w http.ResponseWriter, status int, results []Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
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

	responseMovies(w, 200, results)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	//fmt.Println(movie_id)
	//fmt.Fprintf(w, "Has cargado la pelicula numero %s", movie_id)
	// Comprobar que el ID es un ID hexadecimal
	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)
	//fmt.Println(oid)
	results := Movie{}
	err := collection.FindId(oid).One(&results)
	//fmt.Println(results)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, results)
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

	responseMovie(w, 200, movie_data)
}

func MovieUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(movie_id)

	// Recoger el objeto JSON que nos llega por el BODY
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	// Decodificar los datos y bindearlos en la variable movie_data
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}

	// Cerrar la lectura de algo (body)
	defer r.Body.Close()

	// Hacer el update en la base de datos
	document := bson.M{"_id": oid}
	change := bson.M{"$set": movie_data}
	err = collection.Update(document, change)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, movie_data)
}
