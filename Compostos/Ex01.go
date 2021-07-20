// Existem dois times de futebol, o time amarelo e o time vermelho. O time amarelo tem 5 jogadores (Fernando, João, Lúcia, Mariana e Ana) e o time vermelho tem 4 jogadores (Helena, Jonas, José e Juliana).
// 1) Crie um array de string para cada time e nomeie com o nome do time.
// 2) Printe na tela os nomes dos jogadores de cada time.

package main

import "fmt"

func main() {
	var timeAmarelo [5]string
	timeAmarelo = [5]string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}
	var timeVermelho [4]string
	timeVermelho = [4]string{"Helena", "Jonas", "José", "Juliana"}
	fmt.Println(timeAmarelo)
	fmt.Println(timeVermelho)
}
