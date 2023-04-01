package ejercicios

func Cambio(px, py *int) {
	aux := *px
	*px = *py
	*py = aux
}
