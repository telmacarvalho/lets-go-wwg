// Faça um programa que imprima 10 linhas, onde cada linha obedece a seguinte regra: na linha 1 deve ser impresso um valor (1), na linha 2 devem ser impressos dois valores (1 e 2), na linha três devem ser impressos três valores (1, 2, e 3), … na linha 10 devem ser impressos dez valores (1, 2, 3, 4, 5, 6, 7, 8, 9, 10).

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 1; i++ {
		fmt.Printf("%d\n", i+1)
		for i := 1; i < 2; i++ {
			fmt.Printf("%d%d\n", i, i+1)
			for i := 2; i < 3; i++ {
				fmt.Printf("%d%d%d\n", i-1, i, i+1)
				for i := 3; i < 4; i++ {
					fmt.Printf("%d%d%d%d\n", i-2, i-1, i, i+1)
					for i := 4; i < 5; i++ {
						fmt.Printf("%d%d%d%d%d\n", i-3, i-2, i-1, i, i+1)
						for i := 5; i < 6; i++ {
							fmt.Printf("%d%d%d%d%d%d\n", i-4, i-3, i-2, i-1, i, i+1)
							for i := 6; i < 7; i++ {
								fmt.Printf("%d%d%d%d%d%d%d\n", i-5, i-4, i-3, i-2, i-1, i, i+1)
								for i := 7; i < 8; i++ {
									fmt.Printf("%d%d%d%d%d%d%d%d\n", i-6, i-5, i-4, i-3, i-2, i-1, i, i+1)
									for i := 8; i < 9; i++ {
										fmt.Printf("%d%d%d%d%d%d%d%d%d\n", i-7, i-6, i-5, i-4, i-3, i-2, i-1, i, i+1)
										for i := 9; i < 10; i++ {
											fmt.Printf("%d%d%d%d%d%d%d%d%d%d\n", i-8, i-7, i-6, i-5, i-4, i-3, i-2, i-1, i, i+1)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
