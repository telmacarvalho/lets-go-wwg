// No Exercício #06 da seção "Exercícios", usamos for range para percorrer uma slice de string que representava uma lista de itens a comprar no mercado. Agora, resolva o mesmo exercício usando a sintaxe básica da instrução for (sintaxe apresentada aqui).

package main

import "fmt"

func main() {
	var lista = []string{"ovos", "leite", "granola", "iogurte"}
	for i := 0; i < 4; i++ {
		fmt.Printf("%d - %s\n", i+1, lista[i])
	}
}
