package main

import (
	"estructuras_interfaces/figuras"
	"fmt"
)

func main() {
	var opcion string
	arreglo_figuras := []figuras.Figura{}

	for condicion := true; condicion; condicion = len(arreglo_figuras) != 5 {

		figuras_restantes := 5 - len(arreglo_figuras)
		fmt.Printf("Seleccione un figura ingresando la letra de la opcion hasta completar 5 figuras, figuras restantes: %v", figuras_restantes)
		fmt.Println("\n- a. Rectangulo\n- b. Cuadrado\n- c. Circulo")
		fmt.Scan(&opcion)

		x := 0
		y := 0
		lado := 0
		radio := 1

		switch opcion {
		case "a":
			fmt.Println("Inserte la cordenada x del Punto 1")
			fmt.Scan(&x)
			fmt.Println("Inserte la coordenada y del Punto 1")
			fmt.Scan(&y)
			p1 := figuras.Punto{x, y}
			fmt.Println("Inserte la cordenada x del Punto 2")
			fmt.Scan(&x)
			fmt.Println("Inserte la coordenada y del Punto 2")
			fmt.Scan(&y)
			p2 := figuras.Punto{x, y}
			r := figuras.Rectangulo{P1: p1, P2: p2}
			arreglo_figuras = append(arreglo_figuras, r)
		case "b":
			fmt.Println("Inserte la cordenada x del Punto de 1")
			fmt.Scan(&x)
			fmt.Println("Inserte la coordenada y del Punto 1")
			fmt.Scan(&y)
			p1 := figuras.Punto{x, y}
			fmt.Println("Inserte el tamaño del Lado")
			fmt.Scan(&lado)
			c := figuras.Cuadrado{Pto: p1, Lado: lado}
			arreglo_figuras = append(arreglo_figuras, c)
		case "c":
			fmt.Println("Inserte la cordenada x del Punto de 1")
			fmt.Scan(&x)
			fmt.Println("Inserte la coordenada y del Punto 1")
			fmt.Scan(&y)
			p1 := figuras.Punto{x, y}
			fmt.Println("Inserte el tamaño del Radio")
			fmt.Scan(&radio)
			cir := figuras.Circulo{Pto: p1, Radio: radio}
			arreglo_figuras = append(arreglo_figuras, cir)
		default:
			fmt.Println("Opcion seleccionada no valida")
		}

	}

	for _, f := range arreglo_figuras {
		fmt.Println(f.ToString())
	}
}
