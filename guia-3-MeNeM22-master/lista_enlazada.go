package guia3

import "github.com/untref-ayp2/data-structures/types"

// ListaEnlazada es un lista enlazada simple con nodo ficticio
type ListaEnlazada[T types.Ordered] struct {
	cabeza  *Nodo[T]
	cola    *Nodo[T]
	tamanio int
}

// NuevaLista crea una nueva lista enlazada
func NuevaLista[T types.Ordered]() *ListaEnlazada[T] {
	var dato T
	n1 := NuevoNodo[T](dato)
	n2 := NuevoNodo[T](dato)
	n1.SetNext(n2)
	n2.SetPrev(n1)

	return &ListaEnlazada[T]{cabeza: n1, cola: n2}
}

// Append agrega un nuevo nodo al final de la lista
func (l *ListaEnlazada[T]) Append(dato T) {
	l.cola.InsertBefore(dato)
	l.tamanio++
}

// Prepend agrega un nuevo nodo al principio de la lista
func (l *ListaEnlazada[T]) Prepend(dato T) {
	l.cabeza.InsertAfter(dato)
	l.tamanio++
}

// RemoveFirst elimina el primer nodo de la lista
func (l *ListaEnlazada[T]) RemoveFirst() {
	if l.IsEmpty() {
		return
	}
	l.cabeza.RemoveNext()
	l.tamanio--
}

// RemoveLast elimina el último nodo de la lista
func (l *ListaEnlazada[T]) RemoveLast() {
	if l.IsEmpty() {
		return
	}
	l.cola.Prev().Remove()
	l.tamanio--
}

// Remove elimina el nodo con el dato pasado por parámetro
func (l *ListaEnlazada[T]) Remove(dato T) {
	nodo := l.cabeza.Next()
	for nodo != l.cola {
		if nodo.Data() == dato {
			nodo.Remove()
			l.tamanio--
			return
		}
		nodo = nodo.Next()
	}
}

// IsEmpty devuelve true si la lista está vacía
func (l *ListaEnlazada[T]) IsEmpty() bool {
	return l.tamanio == 0
}

// Size devuelve la cantidad de nodos en la lista
func (l *ListaEnlazada[T]) Size() int {
	return l.tamanio
}

// Head devuelve el dato del primer nodo de la lista
func (l *ListaEnlazada[T]) Head() *Nodo[T] {
	if l.IsEmpty() {
		return nil
	}
	return l.cabeza.Next()
}

// Tail devuelve el dato del último nodo de la lista
func (l *ListaEnlazada[T]) Tail() *Nodo[T] {
	if l.IsEmpty() {
		return nil
	}
	return l.cola.Prev()
}

// Find busca un nodo con el dato pasado por parámetro y lo devuelve
func (l *ListaEnlazada[T]) Find(dato T) *Nodo[T] {
	nodo := l.cabeza.Next()
	for nodo != l.cola {
		if nodo.Data() == dato {
			return nodo
		}
		nodo = nodo.Next()
	}
	return nil
}

// Clear elimina todos los nodos de la lista
func (l *ListaEnlazada[T]) Clear() {
	l.cabeza.SetNext(nil)
	l.cola = l.cabeza
	l.tamanio = 0
}
