package guia3

import "testing"

func TestSortedList(t *testing.T) {
	sl := NewList[int]()

	sl.Insert(3)
	sl.Insert(1)
	sl.Insert(2)
	sl.Insert(-4)
	sl.Insert(3)
	sl.Insert(0)
	sl.Insert(5)

	nodo := sl.lista.Head()
	for i, v := range []int{-4, 0, 1, 2, 3, 3, 5} {
		if nodo.Data() != v {
			t.Errorf("Error en el nodo %d", i)
		}
		nodo = nodo.Next()
	}

}
