package main

import	"fmt"

func reverter(arr []int, start, end int) {
	if start >= end {
		return
	}
	
	arr[start], arr[end] = arr[end], arr[start]
	reverter(arr, start+1, end-1)
}

func main() {
    
	var n int

	fmt.Print("Digite o tamanho do array: ")
	fmt.Scan(&n)

	arr1 := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Digite o valor %d: ", i+1)
		fmt.Scan(&arr1[i])
	}

	reverter(arr1, 0, n-1)

	fmt.Println("Array invertido:", arr1)
}
