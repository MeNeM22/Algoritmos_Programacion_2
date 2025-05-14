package guia3

import (
	"errors"

	dl "github.com/untref-ayp2/data-structures/lists/double_linked_list"
	"github.com/untref-ayp2/data-structures/types"
)

type Cola[T types.Ordered] struct {
	lista *dl.DoubleLinkedList[T]
}

func NewCola[T types.Ordered]() *Cola[T] {
	return &Cola[T]{dl.NewList[T]()}
}

func (c *Cola[T]) Enqueue(dato T) {
	c.lista.Prepend(dato)
}

func (c *Cola[T]) Dequeue() (T, error) {
	var aux T
	if c.lista.IsEmpty() {
		return aux, errors.New("clave inexistente")
	}
	aux = c.lista.Tail().Data()
	c.lista.RemoveLast()
	return aux, nil
}

func (c *Cola[T]) Front() (T, error) {
	aux := c.lista.Head().Data()
	if c.lista.IsEmpty() {
		return aux, errors.New("clave inexistente")
	}
	return aux, nil
}

func (c *Cola[T]) IsEmpty() bool {
	return c.lista.IsEmpty()
}
