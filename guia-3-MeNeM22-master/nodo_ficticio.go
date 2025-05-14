package guia3

import "github.com/untref-ayp2/data-structures/types"

type Nodo[T types.Ordered] struct {
	dato T
	sig  *Nodo[T]
	prev *Nodo[T]
}

func NuevoNodo[T types.Ordered](dato T) *Nodo[T] {
	return &Nodo[T]{dato: dato}
}

func (n *Nodo[T]) Data() T {
	return n.dato
}

func (n *Nodo[T]) Next() *Nodo[T] {
	return n.sig
}

func (n *Nodo[T]) Prev() *Nodo[T] {
	// Implementar
	return n.prev
}

func (n *Nodo[T]) SetData(dato T) {
	n.dato = dato
}

func (n *Nodo[T]) SetNext(sig *Nodo[T]) {
	n.sig = sig
}

func (n *Nodo[T]) SetPrev(prev *Nodo[T]) {
	n.prev = prev
}

func (n *Nodo[T]) InsertAfter(dato T) {
	newNode := NuevoNodo(dato)
	newNode.sig = n.sig
	newNode.prev = n
	n.sig.prev = newNode
	n.sig = newNode
}

func (n *Nodo[T]) InsertBefore(dato T) {
	newNode := NuevoNodo(dato)
	newNode.prev = n.prev
	newNode.sig = n
	n.prev.sig = newNode
	n.prev = newNode

}

func (n *Nodo[T]) Remove() {
	n.prev.sig = n.sig //[N1] <--> [N] <--> [N2]
	n.sig.prev = n.prev
}

func (n *Nodo[T]) RemoveNext() {
	n.sig.Remove() //[N1] <--> [N] <--> [Nremove] <--> [N2]
}

func (n *Nodo[T]) RemovePrev() {
	n.prev.Remove() //[N1] <--> [Nremove] <--> [N] <--> [N2]
}

/*type Nodo[T types.Ordered] struct {
	dato T
	sig  *Nodo[T]
	prev *Nodo[T]
}

func NuevoNodo[T types.Ordered](dato T) *Nodo[T] {
	return &Nodo[T]{dato: dato}
}

func (n *Nodo[T]) Data() T {
	return n.dato
}

func (n *Nodo[T]) Next() *Nodo[T] {
	return n.sig
}

func (n *Nodo[T]) Prev() *Nodo[T] {
	// Implementar
	return n.prev
}

func (n *Nodo[T]) SetData(dato T) {
	n.dato = dato
}

func (n *Nodo[T]) SetNext(sig *Nodo[T]) {
	n.sig = sig
}

func (n *Nodo[T]) SetPrev(prev *Nodo[T]) {
	n.prev = prev
}

func (n *Nodo[T]) InsertAfter(dato T) {
	newnodo := NuevoNodo(dato) //   [NewNodo]
	newnodo.sig = n.sig        // [N] <--> [N2]
	newnodo.prev = n
	n.sig.prev = newnodo
	n.sig = newnodo
}

func (n *Nodo[T]) InsertBefore(dato T) {
	newnodo := NuevoNodo(dato) //   [NewNodo]
	newnodo.sig = n            // [N1] <--> [N]
	newnodo.prev = n.prev
	n.prev.sig = newnodo
	n.prev = newnodo
}

func (n *Nodo[T]) Remove() {
	n.prev.sig = n.sig //[N1] <--> [N] <--> [N2]
	n.sig.prev = n.prev
}

func (n *Nodo[T]) RemoveNext() {
	n.sig.Remove() //[N1] <--> [N] <--> [Nremove] <--> [N2]
}

func (n *Nodo[T]) RemovePrev() {
	n.prev.Remove() //[N1] <--> [Nremove] <--> [N] <--> [N2]
}
*/
