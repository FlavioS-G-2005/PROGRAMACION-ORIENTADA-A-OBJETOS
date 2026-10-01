package main

import (
	"bufio"
	"fmt"
	"os"
	"taller/utilidades"
)

func main() {
	var dolares float64
	var moneda string

	// --- PRUEBA DEL CONVERSOR ---
	fmt.Println("Ingrese la cantidad en dolares:")
	fmt.Scanln(&dolares)

	fmt.Println("Ingrese la moneda (Euros, LB, Won, BTC):")
	fmt.Scanln(&moneda)

	utilidades.Convertir(dolares, moneda)

	// --- PRUEBA DEL CONTADOR ---
	fmt.Println("\nIngrese una moneda solamente):")
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()

	utilidades.Contar(frase)
}