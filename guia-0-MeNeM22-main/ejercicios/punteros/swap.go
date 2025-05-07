package punteros

// Escribir una función con la siguiente firma func Swap(px, py *int) que permita intercambiar el valor almacenado en dos variables enteras a través de punteros.
// Pre: px y py son punteros a enteros.
// Post: Los valores de las variables apuntadas por px y py han sido intercambiados.
func Swap(px, py *int) {
	x := *px
	*px = *py
	*py = x
}
