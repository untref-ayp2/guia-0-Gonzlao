package main

import (
	"fmt"
	"punteros/ejercicios"
)

func main() {
	//Ejercicio 1
	x := 9
	y := 5

	ejercicios.Cambio(&x, &y)

	fmt.Printf("x = %v y = %v\n", x, y)
}
