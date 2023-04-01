package main

import (
	"arreglos/ejercicios"
	"fmt"
)

func main() {
	//Ejercicio 1
	numeros := []int{20, 4, 6, 8, 14, 5}
	resultado := ejercicios.Suma(numeros)
	fmt.Println("1#")
	fmt.Println("La suma es:", resultado)

	//Ejercicio 2
	vector1 := []int{9, 2, 6, 1, 8, 4, 7, 3, 5}
	vector2 := []int{19, 2, 6, 13, 8, 4, 7, 23, 5}
	suma_vectores, producto := ejercicios.SumaProductoVectores(vector1, vector2)
	fmt.Println("\n2#")
	fmt.Println("La suma de los vectores es:", suma_vectores)
	fmt.Println("El producto escalar es:", producto)

	//Ejercicio 3
	conjunto1 := []int{9, 2, 6, 1, 8, 4, 7, 3, 5}
	conjunto2 := []int{19, 2, 6, 13, 8, 4, 23}
	union, interseccion := ejercicios.UnionInterseccionConjuntos(conjunto1, conjunto2)
	fmt.Println("\n3#")
	fmt.Println("La union de los conjuntos es:", union)
	fmt.Println("La interseccion de los conjuntos es:", interseccion)
}
