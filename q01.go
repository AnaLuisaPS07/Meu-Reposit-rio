package main

import	"fmt"

func elevado(x, n int) int {
	if n == 0 {
		return 1
	}

	return x * elevado(x, n-1)
}

func main() {
	
	var b, e int

	fmt.Print("Digite a base: ")
	fmt.Scan(&b)

	fmt.Print("Digite o expoente: ")
	fmt.Scan(&e)

	resultado := elevado(b, e)
	
	fmt.Printf("O resultado de %d^%d é igual a %d \n", b, e, resultado)
}
