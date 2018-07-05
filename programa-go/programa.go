package main

import (
	"fmt"
)

func main() {
	fmt.Println("**** MI PROGRAMA CON GO ****")

	/*
		fmt.Println("Hola " + os.Args[1] + " bienvenido al programa en GO")

		// Condicional if
		edad, _ := strconv.Atoi(os.Args[2])

		if edad >= 18 || edad == 17 {
			fmt.Println("Eres mayor de edad o tienes 17")
		} else {
			fmt.Println("Eres MENOR de edad o demasiado mayor")
		}
	*/

	// Bucle FOR
	peliculas := []string{"Harry Potter", "Pokemon", "Saint Seiya"}
	/*
		for i := 0; i < len(peliculas); i++ {
			if i%2 == 0 {
				fmt.Println("La pelicula "+peliculas[i]+" es par: ", i)
			} else {
				fmt.Println("La pelicula "+peliculas[i]+" es IMPAR: ", i)
			}
		}*/

	// FOREACH
	for _, pelicula := range peliculas {
		fmt.Println(pelicula)
	}

}
