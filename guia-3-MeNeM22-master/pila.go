package guia3

import (
	"errors"

	dl "github.com/untref-ayp2/data-structures/lists/single_linked_list"
	"github.com/untref-ayp2/data-structures/types"
)

type Pila[T types.Ordered] struct {
	lista *dl.SingleLinkedList[T]
}

func NewPila[T types.Ordered]() *Pila[T] {
	return &Pila[T]{lista: dl.NewList[T]()}
}

func (p *Pila[T]) Push(dato T) {
	p.lista.Append(dato)
}

func (p *Pila[T]) Pop() (T, error) {
	var aux T
	if p.lista.IsEmpty() {
		return aux, errors.New("nop")
	}
	aux = p.lista.Tail().Data()
	p.lista.RemoveLast()
	return aux, nil
}

func (p *Pila[T]) Top() (T, error) {
	var aux T
	if p.lista.IsEmpty() {
		return aux, errors.New("nop")
	}
	aux = p.lista.Tail().Data()
	return aux, nil
}

func (p *Pila[T]) IsEmpty() bool {
	return p.lista.IsEmpty()
}
