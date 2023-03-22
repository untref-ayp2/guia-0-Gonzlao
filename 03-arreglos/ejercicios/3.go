package ejercicios

import (
	"sort"
)

func UnionInterseccionConjuntos(conjunto1, conjunto2 []int) ([]int, []int) {
	union := []int{}
	interseccion := []int{}

	valores := make(map[int]bool)

	for _, valor := range conjunto1 {
		valores[valor] = true
	}

	for _, valor := range conjunto2 {
		valores[valor] = true
	}

	for valor := range valores {
		union = append(union, valor)
	}

	for i := 0; i < len(conjunto1); i++ {
		for j := 0; j < len(conjunto2); j++ {
			if conjunto1[i] == conjunto2[j] {
				interseccion = append(interseccion, conjunto1[i])
			}
		}
	}

	sort.Ints(union)
	sort.Ints(interseccion)

	return union, interseccion
}
