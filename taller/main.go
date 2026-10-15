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

	fmt.Println("Ingrese la cantidad en dolares:")
	fmt.Scanln(&dolares)

	fmt.Println("Ingrese la moneda (Euros, LB, Won, BTC):")
	fmt.Scanln(&moneda)

	utilidades.Convertir(dolares, moneda)

	 
	fmt.Println("\nIngrese una frase para contar sus vocales):")
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()

	utilidades.Contar(frase)
}