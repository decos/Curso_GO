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

	var numero1 float32 = 10
	var numero2 float32 = 6

	fmt.Println("Calculadora 1")
	calculadora(numero1, numero2)
	fmt.Println("-------------")

	var numero3 float32 = 44
	var numero4 float32 = 7

	fmt.Println("Calculadora 2")
	calculadora(numero3, numero4)

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
	holaMundo()

	// Retorno de datos
	fmt.Println(devolverTexto())
}

func devolverTexto() (dato1 string, dato2 int) {
	dato1 = "Diego"
	dato2 = 27

	return
}

func holaMundo() {
	fmt.Println("Hola mundo!")
}

func operacion(n1 float32, n2 float32, op string) float32 {
	var resultado float32
	if op == "+" {
		resultado = n1 + n2
	}

	if op == "-" {
		resultado = n1 - n2
	}

	if op == "/" {
		resultado = n1 / n2
	}

	if op == "*" {
		resultado = n1 * n2
	}

	return resultado
}

func calculadora(numero1 float32, numero2 float32) {
	//Suma
	fmt.Print("La suma es: ")
	fmt.Println(operacion(numero1, numero2, "+"))

	//Resta
	fmt.Print("La resta es: ")
	fmt.Println(operacion(numero1, numero2, "-"))

	//Multiplicacion
	fmt.Print("La multiplicacion es: ")
	fmt.Println(operacion(numero1, numero2, "*"))

	//Division
	fmt.Print("La division es: ")
	fmt.Println(operacion(numero1, numero2, "/"))
}
