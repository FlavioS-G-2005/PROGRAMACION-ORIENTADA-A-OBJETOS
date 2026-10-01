package utilidades

import (
	"fmt"
	"strings"
)

func Contar(frase string) {
	a := 0
	e := 0
	i := 0
	o := 0
	u := 0

	// Convertimos todo a minúsculas
	frase = strings.ToLower(frase)

	for j := 0; j < len(frase); j++ {
		letra := frase[j]

		if letra == 'a' {
			a++
		}
		if letra == 'e' {
			e++
		}
		if letra == 'i' {
			i++
		}
		if letra == 'o' {
			o++
		}
		if letra == 'u' {
			u++
		}
	}

	fmt.Println("A:", a)
	fmt.Println("E:", e)
	fmt.Println("I:", i)
	fmt.Println("O:", o)
	fmt.Println("U:", u)
}