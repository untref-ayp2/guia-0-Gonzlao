package ejercicios

func SumaProductoVectores(vector1, vector2 []int) ([]int, int) {
	suma_vectores := []int{}
	producto := 0
	for i := 0; i < len(vector1); i++ {
		suma := vector1[i] + vector2[i]
		suma_vectores = append(suma_vectores, suma)

		producto += (vector1[i] * vector2[i])
	}
	return suma_vectores, producto

}
