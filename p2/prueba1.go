package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numAviones       = 5   // Número de aviones en la simulación
	numPistas        = 2   // Número de pistas en el aeropuerto
	sims             = 3   // Número de simulaciones
	maxEsperaTorre   = 3   // Máximo número de aviones en espera en la torre de control
	maxEsperaPista   = 2   // Máximo número de aviones en espera en la pista
	tiempoBaseTorre  = 1   // Tiempo base de uso de la torre de control en segundos
	tiempoBasePista  = 2   // Tiempo base de uso de la pista en segundos
	tiempoBasePuerta = 3   // Tiempo base de uso de la puerta de desembarque en segundos
	variacion        = 0.3 // Variación del tiempo de uso en 30%
)

var wg sync.WaitGroup

// Función para la torre de control
func torreDeControl(solicitudes <-chan int, canalPista chan<- int) {
	for avionID := range solicitudes {
		fmt.Printf("Torre de Control: Avión %d solicitando permiso para aterrizar.\n", avionID)
		// Tiempo variable de la torre de control
		tiempoTorre := tiempoVariable(tiempoBaseTorre)
		time.Sleep(tiempoTorre)
		fmt.Printf("Torre de Control: Permiso concedido al avión %d para aterrizar.\n", avionID)
		canalPista <- avionID // Enviar permiso a la pista
	}
	close(canalPista) // Cerrar el canal de la pista cuando todas las solicitudes han sido procesadas
}

// Función para una pista de aterrizaje
func procesarPista(aterrizajes <-chan int, canalPuertas chan<- int, pistaID int) {
	for avionID := range aterrizajes {
		fmt.Printf("Pista %d: Avión %d está aterrizando.\n", pistaID, avionID)
		// Tiempo variable de aterrizaje en la pista
		tiempoPista := tiempoVariable(tiempoBasePista)
		time.Sleep(tiempoPista)
		fmt.Printf("Pista %d: Avión %d ha aterrizado.\n", pistaID, avionID)
		canalPuertas <- avionID // Enviar a desembarque
	}
}

// Función para manejar el desembarque de los aviones (NO es una goroutine)
func procesarPuertas(desembarques <-chan int) {
	// Procesar los aviones de manera secuencial
	for avionID := range desembarques {
		fmt.Printf("Puerta de Desembarque: Avión %d está desembarcando.\n", avionID)
		// Tiempo variable de desembarque
		tiempoPuerta := tiempoVariable(tiempoBasePuerta)
		time.Sleep(tiempoPuerta)
		fmt.Printf("Puerta de Desembarque: Desembarco del avión %d completado.\n", avionID)
		wg.Done() // Indicar que el avión ha completado su ciclo
	}
}

// Función para calcular un tiempo variable con variación aleatoria
func tiempoVariable(tiempoBase int) time.Duration {
	variacionFactor := 1 + ((rand.Float64()*2)-1)*variacion
	return time.Duration(float64(tiempoBase)*variacionFactor) * time.Second
}

func ejecutarSimulacion() {
	// Canales para cada etapa del flujo de aviones con límite de espera
	solicitudes := make(chan int, maxEsperaTorre) // Solicitudes de los aviones a torre de control
	canalPista := make(chan int, maxEsperaPista)  // Canal de torre de control a pista
	canalPuertas := make(chan int, numAviones)    // Canal de pista a puerta de desembarque

	// Iniciar goroutine para la torre de control
	go torreDeControl(solicitudes, canalPista)

	// Iniciar múltiples goroutines para pistas
	for i := 1; i <= numPistas; i++ {
		go procesarPista(canalPista, canalPuertas, i)
	}

	// Enviar aviones (IDs) a la torre de control
	for i := 1; i <= numAviones; i++ {
		wg.Add(1)        // Añadir una tarea al WaitGroup para cada avión
		solicitudes <- i // Enviar solicitud de aterrizaje del avión
	}
	close(solicitudes) // Cerramos el canal de solicitudes tras encolar todos los aviones

	// Ejecutar la función secuencialmente para el desembarque de aviones
	procesarPuertas(canalPuertas)

	// Esperar a que todos los aviones completen su desembarque
	wg.Wait()
	fmt.Println("Simulación completada: Todos los aviones han aterrizado y desembarcado.")
}

func main() {
	// Inicializar generador de números aleatorios
	rand.Seed(time.Now().UnixNano())

	// Ejecutar la simulación múltiples veces
	for i := 1; i <= sims; i++ {
		fmt.Printf("\n=== Ejecución de la Simulación %d ===\n", i)
		ejecutarSimulacion()
	}
}
