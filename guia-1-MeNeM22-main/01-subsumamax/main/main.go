package main

import (
	"fmt"
	"subsumamax"
)

func main() {
	arreglo := []int{-1, 6, -2, 5, -1, 4, 3, -4, 3}
	sumaMaxima, posInicial, posFinal := subsumamax.SubsecuenciaSumaMaxima(arreglo)
	fmt.Printf("La subsecuencia con suma máxima es: %v\n", arreglo[posInicial:posFinal+1])
	fmt.Printf("La suma máxima es: %v\n", sumaMaxima)
	fmt.Printf("La posición inicial es: %v\n", posInicial)
	fmt.Printf("La posición final es: %v\n", posFinal)

}
