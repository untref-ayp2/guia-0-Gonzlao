package ejercicios

func GetMaxMin(numeros []int) (int, int) {
	min := numeros[0]
	max := numeros[0]
	for i := 0; i < len(numeros); i++ {
		if numeros[i] < min {
			min = numeros[i]
		}
		if numeros[i] > max {
			max = numeros[i]
		}
	}
	return min, max
}
