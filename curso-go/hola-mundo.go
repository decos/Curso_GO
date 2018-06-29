package main

import (
	"fmt"
)

// Gorra : Tipo de dato personalizado
type Gorra struct {
	marca  string
	color  string
	precio float32
	plana  bool
}

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

	numero1 := 10.
	numero2 := 6.

	//Division
	fmt.Print("La division es: ")
	fmt.Println(numero1 / numero2)

	// Importar tipo de dato personalizado
	/*
		var gorraNegra = Gorra{
			marca:  "Nike",
			color:  "Negro",
			precio: 25.05,
			plana:  false}
	*/

	var gorraNegra = Gorra{"Adidas", "Roja", 25.05, false}

	//fmt.Println(gorraNegra)
	fmt.Println(gorraNegra.marca)
}
