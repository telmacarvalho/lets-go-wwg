// Construa uma função que receba uma palavra ou frase e uma letra, e retorne o número de ocorrências da letra informada.

package main

import "fmt"

func main() {
	letra, quantidade := Letra("Gosto da letra t")
	fmt.Printf("Número de ocorrências da letra %s é %d.\n", letra, quantidade)
}

func Letra(frase string) (string, int) {
	var letra string
	var quantidade int
	ultimaLetra := len([]rune(frase))
	letra = frase[(ultimaLetra - 1):ultimaLetra]
	for _, v := range frase {
		if string(v) == letra {
			quantidade += 1
		}
	}
	return letra, quantidade
}
