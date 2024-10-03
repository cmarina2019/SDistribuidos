package main

import "fmt"

func max(numeros []int) int {
	max := numeros[0]
	for _, num := range numeros {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	numeros := []int{10, 20, 5, 30, 15}
	fmt.Println(max(numeros))
}