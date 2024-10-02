package main

import "fmt"

// Función para inicializar el array con valores
func inicArray(a *[5]int) {
	for i := 0; i < len(a); i++ {
		a[i] = i * 10 // se asigna un valor
	}
}

func main() {

	var a [5]int //declarar un array de tamaño 5
	fmt.Println("Antes de inicializar: ", a)

	inicArray(&a) // inicializar el array
	fmt.Println("Despues de inicializar: ", a)

	fmt.Println("Posiciones y valores del array")
	for i := 0; i < len(a); i++{
		fmt.Printf("Posicion %d: %d\n", i, a[i])
	}

}
