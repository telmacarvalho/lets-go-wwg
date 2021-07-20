// Escreva um programa que lista o nome dos países e quantas vezes eles aparecem no nosso map.
// 1) Criar um mapa com chave do tipo string e valor do tipo string (map[string]string) onde a chave seja o nome da cidade e o valor o nome do país.
// 2) Percorrer o mapa do item 1 acumulando em outro mapa a frequência de aparições do país. Esse mapa segundo mapa terá tipo map[string]int, sendo a chave o nome do país e o valor a quantidade de menções.
// 3) Printe na tela os valores do último mapa.

package main

import "fmt"

func main() {
	cidades := make(map[string]string)
	cidades["Rio de Janeiro"] = "Brasil"
	cidades["Paris"] = "França"
	cidades["Londres"] = "Inglaterra"
	cidades["São Paulo"] = "Brasil"
	cidades["Chicago"] = "EUA"
	cidades["Nova York"] = "EUA"
	cidades["Toronto"] = "Canadá"
	cidades["Barra Mansa"] = "Brasil"
	fmt.Println(cidades)

	contagem := make(map[string]int)

	for i, v := range cidades {
		if v == "Brasil" {
			contagem["Brasil"] = contagem["Brasil"] + 1
		} else if v == "França" {
			contagem["França"] = contagem["França"] + 1
		} else if v == "Inglaterra" {
			contagem["Inglaterra"] = contagem["Inglaterra"] + 1
		} else if v == "EUA" {
			contagem["EUA"] = contagem["EUA"] + 1
		} else if v == "Canadá" {
			contagem["Canadá"] = contagem["Canadá"] + 1
		} else {
			fmt.Println(i)
		}
	}

	fmt.Println(contagem)
}
