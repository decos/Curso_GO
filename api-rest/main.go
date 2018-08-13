package main

import (
	"log"
	"net/http"
)

func main() {
	// Implementar el enrutador
	router := NewRouter()

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
