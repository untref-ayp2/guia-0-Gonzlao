package main

import (
	"estructuras_interfaces/figuras"
	"fmt"
)

func mostrar(f figuras.Figura) {
	fmt.Println(f.ToString())
	fmt.Println("Area: ", f.Area())
	fmt.Println("Perimetro: ", f.Perimetro())
}

func main() {
	p1 := figuras.Punto{0, 0}
	p2 := figuras.Punto{10, 5}

	r := figuras.Rectangulo{P1: p1, P2: p2}
	c := figuras.Cuadrado{Pto: p1, Lado: 10}
	cir := figuras.Circulo{Pto: p1, Radio: 9}

	arreglo_figuras := [3]figuras.Figura{r, c, cir}

	for _, f := range arreglo_figuras {
		mostrar(f)
	}
}
