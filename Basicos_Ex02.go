//1) Código:

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

//2)Podemos pegar a informação do ano diretamente do console através do método fmt.scan().
