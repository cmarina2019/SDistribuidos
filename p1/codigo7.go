package main

import "fmt"

// Función para inicializar el array con valores
func inicArray(a *[5]int) {
	for i := 0; i < len(a); i++ {
		a[i] = i * 10 // Se asigna un valor
	}
}

// Función para calcular la media
func media(numeros [5]float64) float64 {
	suma := 0.0
	for _, num := range numeros {
		suma += num
	}
	return suma / float64(len(numeros))
}

// Función para encontrar el valor máximo
func max(numeros [5]int) int {
	max := numeros[0]
	for _, num := range numeros {
		if num > max {
			max = num
		}
	}
	return max
}

// Función para encontrar el valor mínimo
func min(numeros [5]int) int {
	min := numeros[0]
	for _, num := range numeros {
		if num < min {
			min = num
		}
	}
	return min
}

// Función para encontrar la posición de un valor en el array
func posicion(numeros [5]int, valor int) int {
	for i, num := range numeros {
		if num == valor {
			return i
		}
	}
	return -1
}

func main() {
	var a [5]int // Declarar array de tamaño 5
	inicArray(&a) // Inicializar el array
	fmt.Println("Array inicializado:", a)

	for {
		fmt.Println("\n--- Menú ---")
		fmt.Println("1. Calcular media de números")
		fmt.Println("2. Encontrar valor máximo")
		fmt.Println("3. Encontrar valor mínimo")
		fmt.Println("4. Buscar posición de un valor")
		fmt.Println("5. Salir")
		fmt.Print("Elige una opción: ")

		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			numeros := [5]float64{2, 3, 4, 5, 6} // Puedes modificar estos números
			fmt.Println("La media es:", media(numeros))

		case 2:
			fmt.Println("El valor máximo es:", max(a))

		case 3:
			fmt.Println("El valor mínimo es:", min(a))

		case 4:
			var valor int
			fmt.Print("Ingresa el valor a buscar: ")
			fmt.Scanln(&valor)
			pos := posicion(a, valor)
			if pos != -1 {
				fmt.Printf("El valor %d se encuentra en la posición: %d\n", valor, pos)
			} else {
				fmt.Printf("El valor %d no se encontró en el array\n", valor)
			}

		case 5:
			fmt.Println("Saliendo del programa...")
			return

		default:
			fmt.Println("Opción no válida. Intenta de nuevo.")
		}
	}
}
