package queueS

import "github.com/untref-ayp2/data-structures/stack"

// Implementación de una cola genérica utilizando dos pilas
type QueueS[T any] struct {
	pila1 *stack.Stack[T]
	pila2 *stack.Stack[T]
}

// Crea una nueva cola vacía.
func NewQueueS[T any]() *QueueS[T] {
	return &QueueS[T]{pila1: stack.New[T](), pila2: stack.New[T]()}
}

// Agrega un elemento a la cola.
func (q *QueueS[T]) Enqueue(v T) {
	q.pila1.Push(v)
}

// Elimina y devuelve el elemento al frente de la cola.
func (q *QueueS[T]) Dequeue() (T, error) {
	for !q.pila1.IsEmpty() {
		aux, _ := q.pila1.Pop()
		q.pila2.Push(aux)
	}
	resultado, _ := q.pila2.Pop()
	return resultado, nil
}

// Devuelve el elemento al frente de la cola.
func (q *QueueS[T]) Front() (T, error) {
	for !q.pila1.IsEmpty() {
		aux, _ := q.pila1.Pop()
		q.pila2.Push(aux)
	}
	resultado, _ := q.pila2.Top()
	return resultado, nil
}

// Devuelve true si la cola está vacía.
func (q *QueueS[T]) IsEmpty() bool {
	return q.pila1.IsEmpty() && q.pila2.IsEmpty()
}
