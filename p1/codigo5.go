package main

import "fmt"

func min(numeros []int) int {
	min := numeros[0]
	for _, num := range numeros {
		if num < min {
			min = num
		}
	}
	return min
}

func main() {
	numeros := []int{10, 20, 5, 30, 15, 2}
	fmt.Println(min(numeros))
}