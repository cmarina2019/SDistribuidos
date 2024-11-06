package main

import (
	"fmt"
	"time"
)

const (
	numPistas    = 3    // Número de pistas en el aeropuerto
	numPuertas   = 2    // Número de puertas de desembarque
	numAviones   = 5    // Número de aviones en la simulación
	tiempoPista  = 3    // Tiempo nominal de uso de pista en segundos
	tiempoPuerta = 2    // Tiempo nominal de uso de puerta de desembarque en segundos
	variacion    = 0.25 // Variación del tiempo de uso en 25%
)

type Avion struct {
	id int
}

func main() {
	// Canal para manejar los aviones listos para aterrizar
	colaAterrizaje := make(chan Avion, numAviones)

	// Canal para manejar los aviones listos para desembarque
	puertasDesembarque := make(chan Avion, numPuertas)

	// Crear y encolar aviones en la cola de aterrizaje
	for i := 1; i <= numAviones; i++ {
		avion := Avion{id: i}
		colaAterrizaje <- avion
	}
	close(colaAterrizaje) // Cerramos el canal una vez que todos los aviones están en la cola

	go func() {
		for avion := range puertasDesembarque {
			fmt.Printf("Avion %d está en desembarque .\n", avion.id)
			time.Sleep(time.Duration(tiempoPuerta) * time.Second) // Tiempo de desembarque
			fmt.Printf("Desembarco del avión %d completado.\n", avion.id)
		}
	}()

	// Aterrizaje de los aviones cada 2 segundos
	for avion := range colaAterrizaje {
		fmt.Printf("Avión %d está aterrizando.\n", avion.id)
		time.Sleep(2 * time.Second) // Pausa de 2 segundos entre aterrizajes
		puertasDesembarque <- avion // Enviar avion a desembarque
	}
	close(puertasDesembarque) // Cerramos el canal de puertas cuando todos los aviones hayan aterrizado y embarcado

	fmt.Println("Todos los aviones han aterrizado y embarcado.")
}
