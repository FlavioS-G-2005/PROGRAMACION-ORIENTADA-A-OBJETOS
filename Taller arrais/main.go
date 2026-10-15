package main

import "fmt"

func sacarPromedio(notas []float64) float64 {
	suma := 0.0
	for _, n := range notas {
		suma += n
	}
	return suma / float64(len(notas))
}

func mayorYMenor(notas []float64) (float64, float64) {
	max, min := notas[0], notas[0]
	for _, n := range notas {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return max, min
}

func main() {
	// Matriz de 6 alumnos x 4 materias
	notas := [6][4]float64{
		{8.5, 9.0, 7.5, 8.0},
		{9.5, 9.2, 10.0, 9.8},
		{6.0, 5.5, 7.0, 6.5},
		{7.5, 8.0, 8.5, 7.0},
		{10.0, 10.0, 9.5, 9.0},
		{5.0, 6.5, 6.0, 5.5},
	}

	sumaTotal := 0.0
	totalNotas := 0

	fmt.Println("--- Notas de la clase ---")
	for i, fila := range notas {
		// Pasamos la fila como slice
		sliceNotas := fila[:]

		prom := sacarPromedio(sliceNotas)
		max, min := mayorYMenor(sliceNotas)

		fmt.Printf("Alumno %d -> Promedio: %.2f | Mejor: %.2f | Peor: %.2f\n", i+1, prom, max, min)

		for _, n := range sliceNotas {
			sumaTotal += n
			totalNotas++
		}
	}

	fmt.Printf("\nPromedio general de la clase: %.2f\n", sumaTotal/float64(totalNotas))
} 