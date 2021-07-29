// Construa uma função que receba uma lista de números inteiros, modifique essa lista dobrando os números ímpares e divida por dois os pares, e retorne a lista modificada e a soma de todos os elementos da lista.

package main

import "fmt"

func main() {
	resultado, soma, err := Resultado(0, 1, 2, 3, 4, 5, 6)
	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}
	fmt.Printf("Lista modificada: %v\nSoma de todos os elementos da lista: %v\n", resultado, soma)
}

func Resultado(numeros ...float64) ([]float64, float64, error) {
	var resultado []float64
	var soma float64
	for _, v := range numeros {
		if float64(v) != float64(int64(v)) {
			return []float64{}, 0, fmt.Errorf("Número %v não é inteiro e não é aceito nessa operação", v)
		} else if (int64(v) % 2) != 0 {
			resultado = append(resultado, float64(v*v))
		} else {
			resultado = append(resultado, float64(v/2))
		}
		soma += float64(v)
	}
	return resultado, soma, nil
}
