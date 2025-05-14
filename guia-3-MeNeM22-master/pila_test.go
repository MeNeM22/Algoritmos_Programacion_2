package guia3

import "testing"

func TestPushPop(t *testing.T) {
	p := NewPila[int]()
	if p == nil {
		t.Fatal("La pila no puede ser nula")
	}

	p.Push(1)
	p.Push(2)
	p.Push(3)

	v, err := p.Pop()
	if err != nil || v != 3 {
		t.Error("Error en Push")
	}

	v, err = p.Pop()
	if err != nil || v != 2 {
		t.Error("Error en Push")
	}

	v, err = p.Pop()
	if err != nil || v != 1 {
		t.Error("Error en Push")
	}

	_, err = p.Pop()
	if err == nil {
		t.Error("Error en Push")
	}
}
