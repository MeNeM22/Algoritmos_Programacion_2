package busquedas

// BusLineal busca un elemento en un arreglo de enteros usando el algoritmo de búsqueda lineal
func BusLineal(datos []int, buscado int) int {
	for indice, valor := range datos {
		if buscado == valor {
			return indice
		}
	}
	return -1
}
