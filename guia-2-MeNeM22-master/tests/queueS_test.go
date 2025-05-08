package tests

import (
	"guia2/queueS"
	"testing"
)

func TestQueueS(t *testing.T) {
	q := queueS.NewQueueS[int]()
	if q == nil {
		t.Fatal("La cola no puede ser nula")
	}
	if !q.IsEmpty() {
		t.Error("La cola no está vacía")
	}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if q.IsEmpty() {
		t.Error("La cola está vacía")
	}
	x, _ := q.Dequeue()
	if x != 1 {
		t.Error("Se esperaba 1 y se obtuvo", x)
	}
	x, _ = q.Front()
	if x != 2 {
		t.Error("Se esperaba 2 y se obtuvo", x)
	}
	x, _ = q.Dequeue()
	if x != 2 {
		t.Error("Se esperaba 2 y se obtuvo", x)
	}
	x, _ = q.Dequeue()
	if x != 3 {
		t.Error("Se esperaba 3 y se obtuvo", x)
	}
	if !q.IsEmpty() {
		t.Error("La cola no está vacía")
	}
}

func TestQueueS2(t *testing.T) {
	q := queueS.NewQueueS[string]()
	if q == nil {
		t.Fatal("La cola no puede ser nula")
	}
	if !q.IsEmpty() {
		t.Error("La cola no está vacía")
	}
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")
	if q.IsEmpty() {
		t.Error("La cola está vacía")
	}
	x, _ := q.Dequeue()
	if x != "a" {
		t.Error("Se esperaba a y se obtuvo", x)
	}
	x, _ = q.Front()
	if x != "b" {
		t.Error("Se esperaba b y se obtuvo", x)
	}
	x, _ = q.Dequeue()
	if x != "b" {
		t.Error("Se esperaba b y se obtuvo", x)
	}
	x, _ = q.Dequeue()
	if x != "c" {
		t.Error("Se esperaba c y se obtuvo", x)
	}
	if !q.IsEmpty() {
		t.Error("La cola no está vacía")
	}
}

func TestQueueS3(t *testing.T) {
	q := queueS.NewQueueS[bool]()
	if q == nil {
		t.Fatal("La cola no puede ser nula")
	}
	if !q.IsEmpty() {
		t.Error("La cola no está vacía")
	}
	q.Enqueue(true)
	q.Enqueue(false)
	q.Enqueue(true)
	if q.IsEmpty() {
		t.Error("La cola está vacía")
	}
	x, _ := q.Dequeue()
	if x != true {
		t.Error("Se esperaba true y se obtuvo", x)
	}
	x, _ = q.Front()
	if x != false {
		t.Error("Se esperaba false y se obtuvo", x)
	}
	x, _ = q.Dequeue()
	if x != false {
		t.Error("Se esperaba false y se obtuvo", x)
	}
	x, _ = q.Dequeue()
	if x != true {
		t.Error("Se esperaba true y se obtuvo", x)
	}
	if !q.IsEmpty() {
		t.Error("La cola no está vacía")
	}
}
