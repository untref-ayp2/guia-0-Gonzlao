package ejercicios

func EsPrimo(numero int) bool {
	resultado := true
	if numero%2 == 0 {
		resultado = false
	}
	return resultado
}
