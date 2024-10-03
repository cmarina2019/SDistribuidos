package main

import "fmt"

func media(numeros []float64) float64{
	suma := 0.0
	for _, num:= range numeros {
		suma += num
	}
	return suma / float64(len(numeros))
}

func main(){
	numeros := []float64{2,3,4,5,6}
	fmt.Println(media(numeros))
}
