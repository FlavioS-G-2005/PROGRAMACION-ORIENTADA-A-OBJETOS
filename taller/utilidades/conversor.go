package utilidades

import "fmt"

func Convertir(dolares float64, moneda string) {
	if moneda == "Euros" {
		fmt.Println("En Euros es:", dolares*0.92)
	} else if moneda == "LB" {
		fmt.Println("En Libras es:", dolares*0.79)
	} else if moneda == "Won" {
		fmt.Println("En Won es:", dolares*1350.50)
	} else if moneda == "BTC" {
		fmt.Println("En BTC es:", dolares*0.000015)
	} else {
		fmt.Println("Moneda no reconocida.")
	}
}