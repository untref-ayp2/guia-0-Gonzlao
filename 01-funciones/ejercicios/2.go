package ejercicios

import "fmt"

func Menu(opcion int) {

	switch opcion {
	case 1:
		fmt.Println("Selecciono la Opción 1")
	case 2:
		fmt.Println("Selecciono la Opción 2")
	case 3:
		fmt.Println("Selecciono la Opción 3")
	case 4:
		fmt.Println("Selecciono la Opción 4")
	default:
		fmt.Println("Opción no valida")
	}

}
