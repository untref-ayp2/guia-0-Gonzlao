package main

import (
	"fmt"
	"funciones/ejercicios"
)

func main() {

	//Ejercicio 1
	var coeficientes = []float64{3.1, 2.3, 1.4, 5.5}
	polinomio := ejercicios.Polinomio(coeficientes)
	fmt.Println("1#")
	fmt.Println(polinomio)

	//Ejercicio 2
	fmt.Println("\n2#")
	ejercicios.Menu(5)

	//Ejercicio 3
	numeros := []int{3, 6, 8, 9, 1, 20, 2, 7, 4}
	min, max := ejercicios.GetMaxMin(numeros)
	fmt.Println("\n3#")
	fmt.Printf("El minmo es: %v y el maximo es: %v\n", min, max)
}
