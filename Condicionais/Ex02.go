package main

import "fmt"

func main() {
	fmt.Println("Informe um valor:")
	var numero int
	fmt.Scan(&numero)

	if numero == 0 {
		fmt.Printf("O número é igual a zero, não é multiplo de 2 e nem de 3.")
	} else {
		fmt.Printf("O número não é igual a zero, ")
		if (numero % 2) == 0 {
			fmt.Printf("é múltiplo de 2 ")
		} else {
			fmt.Printf("não é múltiplo de 2 ")
		}
		if (numero % 3) == 0 {
			fmt.Printf("e é múltiplo de 3.")
		} else {
			fmt.Printf("e não é múltiplo de 3.")
		}

	}
}
