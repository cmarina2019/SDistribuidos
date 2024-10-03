package main

import "fmt"

func posicion(numeros []int, valor int) int {
	for i, num := range numeros {
		if num == valor {
			return i
		}
	}
	return -1
}

func main() {

	numeros := []int{10, 20, 5, 30, 15}
	valor := 3
	fmt.Println("El valor se encuentra en la posicion: ",posicion(numeros, valor))
}