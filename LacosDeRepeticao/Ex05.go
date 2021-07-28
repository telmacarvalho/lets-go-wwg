// Faça um programa que, dado um texto inserido pelo usuário, itere nesse texto e conte o número de ocorrências de cada letra. Em seguida imprima em ordem alfabética a letra e o número de ocorrências dela no texto informado.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var text string
	fmt.Println("Insira um texto:")
	text = leiaLinha()
	letras := make(map[string]int)

	for _, letra := range text {
		letras[string(letra)] += 1
	}

	letrasEmOrdem := make([]string, 0, len(letras))
	for i := range letras {
		letrasEmOrdem = append(letrasEmOrdem, i)
	}
	sort.Strings(letrasEmOrdem)

	for _, k := range letrasEmOrdem {
		fmt.Printf("%v => %v, ", k, letras[k])
	}
}

func leiaLinha() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}
