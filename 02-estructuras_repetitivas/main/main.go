package main

import (
	"estructuras_repetitivas/ejercicios"
	"fmt"
)

func main() {
	//Ejercicio 1
	factorial := ejercicios.Factorial(5)
	fmt.Println("1#")
	fmt.Println("El factorial es:", factorial)

	//Ejercicio 2
	resultado := ejercicios.Producto(3, 4)
	fmt.Println("\n2#")
	fmt.Println("El producto es:", resultado)

	//Ejercicio 3
	es_primo := ejercicios.EsPrimo(8)
	fmt.Println("\n3#")
	fmt.Println("El numnero es primo?", es_primo)
}
