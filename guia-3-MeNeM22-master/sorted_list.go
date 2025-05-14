package guia3

import (
	dl "github.com/untref-ayp2/data-structures/lists/double_linked_list"
	"github.com/untref-ayp2/data-structures/types"
)

// Lista ordenada de forma ascendente
type SortedList[T types.Ordered] struct {
	lista *dl.DoubleLinkedList[T]
}

func NewList[T types.Ordered]() *SortedList[T] {
	return &SortedList[T]{dl.NewList[T]()}
}

func (s *SortedList[T]) Insert(dato T) {
	switch {
	case s.lista.IsEmpty():
		s.lista.Append(dato)
	case s.lista.Head().Data() > dato:
		s.lista.Prepend(dato)
	case s.lista.Tail().Data() < dato:
		s.lista.Append(dato)
	default:
		nodo := s.lista.Head()
		for nodo.Next().Data() < dato {
			nodo = nodo.Next()
		}
		siguiente := nodo.Next()
		nuevoNodo := dl.NewNode(dato)
		nodo.SetNext(nuevoNodo)
		nuevoNodo.SetNext(siguiente)
		siguiente.SetPrev(nuevoNodo)
		nuevoNodo.SetPrev(nodo)
	}
}

func (s *SortedList[T]) RemoveFirst() {
	s.lista.RemoveFirst()
}

func (s *SortedList[T]) RemoveLast() {
	s.lista.RemoveLast()
}

func (s *SortedList[T]) Remove(data T) {
	s.lista.Remove(data)
}

func (s *SortedList[T]) IsEmpty() bool {
	return s.lista.IsEmpty()
}

func (s *SortedList[T]) Size() int {
	return s.lista.Size()
}

func (s *SortedList[T]) Head() T {
	return s.lista.Head().Data()
}

func (s *SortedList[T]) Tail() T {
	return s.lista.Tail().Data()
}

func (s *SortedList[T]) Find(data T) *dl.Node[T] {
	return s.lista.Find(data)
}

func (s *SortedList[T]) Clear() {
	s.lista.Clear()
}
