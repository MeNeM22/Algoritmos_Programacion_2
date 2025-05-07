package funciones

// MinMax devuelve el menor y el mayor número de una lista de enteros.
// Pre: La lista no debe estar vacía.
// Post: Devuelve el menor y el mayor número de la lista.
func EncontrarMinimoMaximo(lista []int) (int, int) {
	var maximo = 0
	var minimo = 0
	var numalto = lista[len(lista)-1]
	var numbajo = lista[0]

	for i := 0; i < len(lista); i++ {
		if lista[i] > numalto {
			numalto = lista[i]
			maximo = lista[i]
		} else if lista[i] < numbajo {
			numbajo = lista[i]
			minimo = lista[i]
		}
	}

	return minimo, maximo
}
