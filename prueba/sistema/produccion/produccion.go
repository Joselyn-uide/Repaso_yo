package produccion

import "fmt"

func registroProduccion(valorRegistro []float64) (float64, float64) {
	totalprodu := 0.0
	for _, valor := range valorRegistro {
		totalprodu += valor
	}

	promedio := totalprodu / float64(len(valorRegistro))
	return totalprodu, promedio

	if promedio > 70 {
		fmt.Println("El promedio general es bueno")
	} else {
		fmt.Println("Se debe mejorar el promedio general.")
	}
}
