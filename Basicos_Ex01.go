//1)Não compila porque o tipo numérico usado é int8, que contempla valores de -128 a 127, e o valor que está sendo armazenado é 150.
//2)O erro nos indica que o valor armazenado na constante está transbordando o limite que int8 comporta.
//3)Correção:
package main

import "fmt"

func main() {
	var quilometros int16
	quilometros = 150
	fmt.Println(quilometros)
}
