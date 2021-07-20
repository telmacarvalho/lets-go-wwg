// Considerando os times do item anterior, crie uma slice para representar cada um.
//1) Adicione o jogador Luis Inácio no time vermelho usando a função append()
//2) Printe na tela os nomes dos jogadores do time vermelho
package main

import "fmt"

func main() {
	var timeAmarelo = []string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}
	var timeVermelho = []string{"Helena", "Jonas", "José", "Juliana"}
	timeVermelho = append(timeVermelho, "Luis Incio")

	fmt.Println(timeAmarelo)
	fmt.Println(timeVermelho)
}
