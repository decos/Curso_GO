package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("**** MI PROGRAMA CON GO ****")

	fmt.Println("Hola " + os.Args[1] + " bienvenido al programa en GO")

	// Condicional if
	edad, _ := strconv.Atoi(os.Args[2])

	if edad >= 18 || edad == 17 {
		fmt.Println("Eres mayor de edad o tienes 17")
	} else {
		fmt.Println("Eres MENOR de edad o demasiado mayor")
	}

}
