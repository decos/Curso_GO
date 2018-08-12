package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Lector")

	// Write file
	nuevoTexto := os.Args[1] + "\n"

	// Overwrite file
	//nuevoTexto := []byte(os.Args[1])
	//escribir := ioutil.WriteFile("fichero.txt", nuevoTexto, 0777)
	//showError(escribir)

	// Append text to the file
	fichero, err := os.OpenFile("fichero.txt", os.O_APPEND|os.O_WRONLY, 0777)
	showError(err)

	escribir, err := fichero.WriteString(nuevoTexto)
	fmt.Println(escribir)
	showError(err)

	fichero.Close()

	// Read file
	texto, errorDeFichero := ioutil.ReadFile("fichero.txt")
	showError(errorDeFichero)

	fmt.Println(string(texto))
}

func showError(e error) {
	if e != nil {
		panic(e)
	}
}
