// Considerando os tópicos que já aprendemos até agora: slices, structs ,condicionais e laços de repetição, crie um programa que traga as informações sobre apartamentos de um prédio. Passos:
// 1) Crie uma estrutura que representa um apartamento, com campos para representar seu número, o nome da sua proprietária e se tem vaga de garagem
// 2) Reúna as estruturas em uma slice que representa um conjunto de apartamentos
// 3) Printe as informações de cada unidade, separando por linha, usando for range

package main

import "fmt"

func main() {
	type Apartamento struct {
		Numero       int
		Proprietaria string
		Vaga         bool
	}
	apartamento1 := Apartamento{Numero: 101, Proprietaria: "Penny", Vaga: false}
	apartamento2 := Apartamento{Numero: 102, Proprietaria: "Daenerys", Vaga: true}
	apartamento3 := Apartamento{Numero: 103, Proprietaria: "June", Vaga: false}

	var Apartamentos = []Apartamento{apartamento1, apartamento2, apartamento3}

	for i := range Apartamentos {
		if Apartamentos[i].Vaga == true {
			fmt.Printf("Apartamento %d: a proprietária é %s e possui vaga de garagem.\n", Apartamentos[i].Numero, Apartamentos[i].Proprietaria)
		} else {
			fmt.Printf("Apartamento %d: a proprietária é %s e não possui vaga de garagem.\n", Apartamentos[i].Numero, Apartamentos[i].Proprietaria)
		}
	}

}
