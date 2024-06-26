package main

import "fmt"

func soma(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}

	return arr[0] + soma(arr[1:])
}

func main() {

	var n int

	fmt.Print("Digite o tamanho do array: ")
	fmt.Scan(&n)

	arr1 := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Digite o valor %d: ", i+1)
		fmt.Scan(&arr1[i])
	}

	resultado := soma(arr1)
	fmt.Printf("A soma de todos os valores do array é: %.2f\n", resultado)
}
