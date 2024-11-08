package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numPistas    = 5    // Número de pistas en el aeropuerto
	numPuertas   = 2    // Número de puertas de desembarque
	numAviones   = 2    // Número de aviones en la simulación
	tiempoPista  = 3    // Tiempo nominal de uso de pista en segundos
	tiempoPuerta = 2    // Tiempo nominal de uso de puerta de desembarque en segundos
	variacion    = 0.25 // Variación del tiempo de uso en 25%
)

type Avion struct {
	id int
}

// Funcion de la torre de control que se ejecutará como una goroutine
func torreDeControl(solicAterrizaje <-chan Avion, permAterrizaje chan<- Avion) {
	for Avion := range solicAterrizaje {
		fmt.Printf("Torre de Control: Avion %d solicitando permiso para aterrizar.\n", Avion.id)
		time.Sleep(1 * time.Second) // Tiempo de respuesta
		fmt.Printf("Torre de Control: Permiso concedido para aterrizar al avion %d.\n", Avion.id)
		permAterrizaje <- Avion // Enviar permisos para aterrizar
	}
	close(permAterrizaje) // Cerarr canal de permisos
}

func main() {
	// Canal para manejar los aviones listos para aterrizar
	//colaAterrizaje := make(chan Avion, numAviones)

	// Canal para manejar los aviones listos para desembarque
	puertasDesembarque := make(chan Avion, numPuertas)

	//Canal para comunicacion con la torre de control

	solicAterrizaje := make(chan Avion, numAviones)
	permAterrizaje := make(chan Avion, numPuertas)

	//Iniciar la torre de control como goroutine
	go torreDeControl(solicAterrizaje, permAterrizaje)

	var wg sync.WaitGroup

	// Crear y encolar aviones en la cola de solicitudes
	for i := 1; i <= numAviones; i++ {
		Avion := Avion{id: i}
		solicAterrizaje <- Avion
	}
	close(solicAterrizaje) //Cerrramo el canal una vez todos los aviones han mandado solicitud

	//Goroutine para manejar el desembarque de aviones
	go func() {
		for Avion := range puertasDesembarque {
			fmt.Printf("Avion %d está en desembarque.\n", Avion.id)
			time.Sleep(time.Duration(tiempoPuerta) * time.Second)
			fmt.Printf("Desembarco del avion %d completado.\n", Avion.id)
			wg.Done() //Marcar que el avion ha completado el desembarco
		}
	}()

	//Aterrizaje de los aviones tras recibir permiso de la torre
	for Avion := range permAterrizaje {
		fmt.Printf("Avion %d está aterrizando.\n", Avion.id)
		time.Sleep(2 * time.Second)
		puertasDesembarque <- Avion // Enviar avion a desembarque
		wg.Add(1)                   //Añadir una tarea al WaitGroup
	}
	close(puertasDesembarque) //Cerrar canal de puertas

	wg.Wait() // Esperar a que todos los aviones hayan completado.
	fmt.Println("Todos los aviones han aterrizado y desembarcado.")
}
