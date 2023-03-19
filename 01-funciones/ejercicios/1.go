package ejercicios

import "strconv"

func Polinomio(coeficientes []float64) string {
	polinomio_str := ""
	for i := 0; i < len(coeficientes); i++ {
		potencia := strconv.Itoa(i)
		coeficiente := strconv.FormatFloat(coeficientes[i], 'f', +1, 64)
		var termino string
		if i == 0 {
			termino = coeficiente
		} else if i == 1 {
			termino = termino + " + " + coeficiente + " X"
		} else {
			termino = termino + " + " + coeficiente + " X**" + potencia
		}
		polinomio_str = polinomio_str + termino

	}
	return polinomio_str
}
