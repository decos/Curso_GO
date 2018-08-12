package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fmt.Println("Lector")

	texto, error := ioutil.ReadFile("fichero.txt")
	showError(error)

	fmt.Println(string(texto))
}

func showError(e error) {
	if e != nil {
		panic(e)
	}
}
