// Faça um programa que solicite à usuária que informe um número até que o número informado seja par. Seguindo o fluxo:Insira um número:
// > 3
// Insira um número:
// > 7
// Insira um número:
// > 2
// Agora sim podemos dividir igualmente entre mim e você!
// Em que 3, 7, e 2 são informados pelo usuário.

package main

import (
	"fmt"
	"os"
)

func main() {
	var valor int
	fmt.Printf("Insira um número:\n")
	fmt.Fscanf(os.Stdin, "%d", &valor)
	for (valor % 2) != 0 {
		fmt.Printf("Insira um número:\n")
		fmt.Fscanf(os.Stdin, "%d", &valor)
	}
	fmt.Printf("Agora sim podemos dividir igualmente entre mim e você!")

}
