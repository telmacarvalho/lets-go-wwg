//1) Dado o ano em que a pessoa nasceu, calcule quantos anos ela tem no ano atual (para fins de arredondamento, considere que ela já fez aniversário no ano atual).
//R:
package main

import (
	"fmt"
	"time"
)

func main() {
	tempoAtual := time.Now()
	var anoNascimento int
	fmt.Println("Qual ano você nasceu?")
	fmt.Scan(&anoNascimento)
	idade := tempoAtual.Year() - anoNascimento
	fmt.Println(anoNascimento)
	fmt.Printf("Você tem %v anos.", idade)
}

//2) Como podemos pegar a informação do ano diretamente do console?
//R: Podemos pegar a informação do ano diretamente do console através do método fmt.scan().
