package main

import (
	"fmt"
)

func main() {
	fmt.Println("**** MI PROGRAMA CON GO ****")

	// Condicional if
	edad := 18

	if edad >= 18 || edad == 17 {
		fmt.Println("Eres mayor de edad o tienes 17")
	} else {
		fmt.Println("Eres MENOR de edad o demasiado mayor")
	}
}
