package figuras

import "fmt"

type Circulo struct {
	Pto   Punto
	Radio int
}

func (c Circulo) Mover(incX, incY int) {
	c.Pto.Mover(incX, incY)
}

func (c Circulo) Perimetro() int {
	return 2 * 3 * c.Radio
}

func (c Circulo) Area() int {
	return 3 * c.Radio * c.Radio
}

func (c Circulo) ToString() string {
	cadena := fmt.Sprint("Circulo:", c)

	return cadena
}
