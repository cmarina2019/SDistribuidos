package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDuplicarMaxEsperaTorre(t *testing.T) {
	// Valores originales
	originalMaxEsperaTorre := maxEsperaTorre
	duplicadoMaxEsperaTorre := originalMaxEsperaTorre * 2

	// Simulación con el valor original de maxEsperaTorre
	fmt.Println("\nEjecutando simulación con el valor original de maxEsperaTorre...")
	start := time.Now()
	ejecutarSimulacion(numAviones, numPistas, numPuertas, maxEsperaTorre, maxEsperaPista, tiempoBaseTorre, tiempoBasePista, tiempoBasePuerta, variacion)
	duracionOriginal := time.Since(start)

	// Simulación con el valor duplicado de maxEsperaTorre
	fmt.Println("\nEjecutando simulación con el valor duplicado de maxEsperaTorre...")
	start = time.Now()
	ejecutarSimulacion(numAviones, numPistas, numPuertas, maxEsperaTorre, maxEsperaPista, tiempoBaseTorre, tiempoBasePista, tiempoBasePuerta, variacion)
	duracionDuplicada := time.Since(start)

	// Imprimir y comparar los resultados
	fmt.Printf("\nDuración con maxEsperaTorre original (%d): %v\n", originalMaxEsperaTorre, duracionOriginal)
	fmt.Printf("Duración con maxEsperaTorre duplicado (%d): %v\n", duplicadoMaxEsperaTorre, duracionDuplicada)

}

func TestAumentarVariacion(t *testing.T) {
	// Valores originales de variación
	originalVariacion := variacion
	incrementoVariacion := originalVariacion * 1.25 // Aumentar variación en 25%

	// Simulación con la variación original
	fmt.Println("\nEjecutando simulación con la variación original...")
	start := time.Now()
	ejecutarSimulacion(numAviones, numPistas, numPuertas, maxEsperaTorre, maxEsperaPista, tiempoBaseTorre, tiempoBasePista, tiempoBasePuerta, variacion)
	duracionOriginal := time.Since(start)

	// Simulación con la variación aumentada
	fmt.Println("\nEjecutando simulación con la variación aumentada en un 25%...")
	start = time.Now()
	ejecutarSimulacion(numAviones, numPistas, numPuertas, maxEsperaTorre, maxEsperaPista, tiempoBaseTorre, tiempoBasePista, tiempoBasePuerta, variacion)
	duracionAumentada := time.Since(start)

	// Imprimir y comparar los resultados
	fmt.Printf("\nDuración con variación original (%.2f): %v\n", originalVariacion, duracionOriginal)
	fmt.Printf("Duración con variación aumentada (%.2f): %v\n", incrementoVariacion, duracionAumentada)
}

func TestDuplicarMaxEsperaYVariacion(t *testing.T) {
	// Valores originales
	originalMaxEsperaTorre := maxEsperaTorre
	duplicadoMaxEsperaTorre := originalMaxEsperaTorre * 2
	originalVariacion := variacion
	incrementoVariacion := originalVariacion * 1.25 // Aumentar variación en 25%

	// Simulación con los valores originales
	fmt.Println("\nEjecutando simulación con maxEsperaTorre y variación originales...")
	start := time.Now()
	ejecutarSimulacion(numAviones, numPistas, numPuertas, maxEsperaTorre, maxEsperaPista, tiempoBaseTorre, tiempoBasePista, tiempoBasePuerta, variacion)
	duracionOriginal := time.Since(start)

	// Simulación con el maxEsperaTorre duplicado y variación aumentada
	fmt.Println("\nEjecutando simulación con maxEsperaTorre duplicado y variación aumentada...")
	start = time.Now()
	ejecutarSimulacion(numAviones, numPistas, numPuertas, maxEsperaTorre, maxEsperaPista, tiempoBaseTorre, tiempoBasePista, tiempoBasePuerta, variacion)
	duracionModificada := time.Since(start)

	// Imprimir y comparar los resultados
	fmt.Printf("\nDuración con maxEsperaTorre y variación originales (Max espera: %d, Variación: %.2f): %v\n", originalMaxEsperaTorre, originalVariacion, duracionOriginal)
	fmt.Printf("Duración con maxEsperaTorre duplicado y variación aumentada (Max espera: %d, Variación: %.2f): %v\n", duplicadoMaxEsperaTorre, incrementoVariacion, duracionModificada)
}

func TestMultiplicarPistas(t *testing.T) {
	// Valores originales
	originalNumPistas := numPistas
	numPistasMultiplicadas := originalNumPistas * 5

	// Simulación con el número original de pistas
	fmt.Println("\nEjecutando simulación con el número original de pistas...")
	start := time.Now()
	ejecutarSimulacion(numAviones, numPistas, numPuertas, maxEsperaTorre, maxEsperaPista, tiempoBaseTorre, tiempoBasePista, tiempoBasePuerta, variacion)
	duracionOriginal := time.Since(start)

	// Simulación con el número de pistas multiplicado por 5
	fmt.Println("\nEjecutando simulación con el número de pistas multiplicado por 5...")
	start = time.Now()
	ejecutarSimulacion(numAviones, numPistas, numPuertas, maxEsperaTorre, maxEsperaPista, tiempoBaseTorre, tiempoBasePista, tiempoBasePuerta, variacion)
	duracionMultiplicada := time.Since(start)

	// Imprimir y comparar los resultados
	fmt.Printf("\nDuración con el número original de pistas (%d): %v\n", originalNumPistas, duracionOriginal)
	fmt.Printf("Duración con el número de pistas multiplicado por 5 (%d): %v\n", numPistasMultiplicadas, duracionMultiplicada)
}
