// Faça um programa em que 3 variáveis recebem valores diferentes e informa qual a variável com maior valor.
package main

import "fmt"

func main() {
	valores := make([]float64, 3)
	fmt.Println("Informe um valor:")
	var a float64
	fmt.Scan(&a)
	valores[0] = a
	fmt.Println("Informe outro valor:")
	var b float64
	fmt.Scan(&b)
	valores[1] = b
	fmt.Println("Informe mais um valor:")
	var c float64
	fmt.Scan(&c)
	valores[2] = c

	if valores[0] >= valores[1] && valores[0] >= valores[2] {
		fmt.Printf("O maior valor é %v.", valores[0])
	} else if valores[1] >= valores[0] && valores[1] >= valores[2] {
		fmt.Printf("O maior valor é %v.", valores[1])
	} else if valores[2] >= valores[0] && valores[2] >= valores[1] {
		fmt.Printf("O maior valor é %v.", valores[2])
	}
}
