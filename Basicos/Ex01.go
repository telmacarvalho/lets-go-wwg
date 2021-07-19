//1) Descubra por que não compila
//R: Não compila porque o tipo numérico usado é int8, que contempla valores de -128 a 127, e o valor que está sendo armazenado é 150.

//2)  Erros de compilação nos ajudam a compreender o que precisamos consertar em nosso código. O que o erro ./prog.go:9:14: constant 150 overflows int8 nos indica?
//R: O erro nos indica que o valor armazenado na constante está transbordando o limite que int8 comporta.

//3) Conserte o erro de compilação e faça a quantidade de quilômetros ser imprimida na tela.
//R:
package main

import "fmt"

func main() {
	var quilometros int16
	quilometros = 150
	fmt.Println(quilometros)
}
