package tests

import (
	"fmt"
	"guia2/ejercicios"
	"testing"

	"github.com/untref-ayp2/data-structures/queue"
)

func TestInvertirCadena(t *testing.T) {
	salida := ejercicios.InvertirCadena("abcd")
	fmt.Println(salida)
	if salida != "dcba" {
		t.Error("InvertirCadena falla")
	}
}

func TestInvertirCadenaVacia(t *testing.T) {
	salida := ejercicios.InvertirCadena("")
	if salida != "" {
		t.Error("InvertirCadena falla")
	}
}

func TestInvertirCadenaImpar(t *testing.T) {
	salida := ejercicios.InvertirCadena("abcde")
	if salida != "edcba" {
		t.Error("InvertirCadena falla")
	}
}

func TestPalindromo(t *testing.T) {
	if ejercicios.Palindromo("1456541") != true ||
		ejercicios.Palindromo("145541") != true {
		t.Error("Palindromo falla")
	}
}

func TestPalindromoFalso(t *testing.T) {
	if ejercicios.Palindromo("145654") != false {
		t.Error("Palindromo falla")
	}
}

func TestPalindromoVacio(t *testing.T) {
	if ejercicios.Palindromo("") != true {
		t.Error("Palindromo falla")
	}
}

func TestPalindromoImpar(t *testing.T) {
	if ejercicios.Palindromo("1456541") != true {
		t.Error("Palindromo falla")
	}
}

func TestPalindromoPar(t *testing.T) {
	if true != ejercicios.Palindromo("145541") {
		t.Error("Palindromo falla")
	}
}

func TestBalanceada(t *testing.T) {
	if ejercicios.Balanceada("[()]{}{[()()]()}") != true || ejercicios.Balanceada("[(])") != false {
		t.Error("Balanceada falla")
	}
}

func TestUnirColas(t *testing.T) {
	q1 := queue.New[int]()
	q1.Enqueue(1)
	q1.Enqueue(2)
	q1.Enqueue(3)
	q2 := queue.New[int]()
	q2.Enqueue(5)
	q2.Enqueue(7)
	q12 := ejercicios.UnirColas(q1, q2)
	if q12 == nil {
		t.Fatal("La cola no puede ser nula")
	}

	value, _ := q12.Dequeue()
	if value != 1 {
		t.Error("UnirColas falla")
	}
	value, _ = q12.Dequeue()
	if value != 2 {
		t.Error("UnirColas falla")
	}
	value, _ = q12.Dequeue()
	if value != 3 {
		t.Error("UnirColas falla")
	}
	value, _ = q12.Dequeue()
	if value != 5 {
		t.Error("UnirColas falla")
	}
	value, _ = q12.Dequeue()
	if value != 7 {
		t.Error("UnirColas falla")
	}
	if !q12.IsEmpty() {
		t.Error("UnirColas falla")
	}
}

func TestUnirColasVacias(t *testing.T) {
	q1 := queue.New[int]()
	q2 := queue.New[int]()

	q12_dado := ejercicios.UnirColas(q1, q2)
	if q12_dado == nil {
		t.Fatal("La cola no puede ser nula")
	}
	if !q12_dado.IsEmpty() {
		t.Error("UnirColas falla")
	}
}

func TestUnirColaVacia(t *testing.T) {
	q1 := queue.New[int]()
	q1.Enqueue(1)
	q2 := queue.New[int]()

	q12 := ejercicios.UnirColas(q1, q2)
	if q12 == nil {
		t.Fatal("La cola no puede ser nula")
	}
	value, _ := q12.Dequeue()
	if value != 1 {
		t.Error("UnirColas falla")
	}
	if !q12.IsEmpty() {
		t.Error("UnirColas falla")
	}
}
