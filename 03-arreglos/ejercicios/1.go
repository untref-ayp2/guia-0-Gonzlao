package ejercicios

func Suma(numeros []int) int {
	resultado := 0
	for i := 0; i < len(numeros); i++ {
		resultado += numeros[i]
	}
	return resultado
}
