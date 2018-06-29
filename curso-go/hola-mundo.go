package main

import (
	"fmt"
)

func main() {
	//time.Sleep(time.Second * 5)

	user := "Diego Abanto"
	pais := "Rusia"

	var suma = 8 + 9
	var resta = 6 - 4
	var nombre = "Diego "
	var apellidos = "Abanto Arroyo "
	var prueba = true
	var flotante = 12.34

	const year = 2018

	fmt.Println(year)
	fmt.Println(prueba)
	fmt.Println(flotante)
	fmt.Println("Hola Mundo desde Go con", user)
	fmt.Println("Hola Mundo desde GO con " + nombre + apellidos + pais)
	fmt.Println(suma)
	fmt.Println(resta)
}
