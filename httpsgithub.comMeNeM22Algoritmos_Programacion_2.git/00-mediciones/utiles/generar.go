package utiles

import (
	"math/rand"
	"time"
)

// GenerarArreglo genera un arreglo de tamaño tam con números aleatorios entre 0 y numeros
func GenerarArreglo(numeros, tam int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arreglo := make([]int, tam)
	for i := 0; i < tam; i++ {
		arreglo[i] = r.Intn(numeros)
	}
	return arreglo
}
