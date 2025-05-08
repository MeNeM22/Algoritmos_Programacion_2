package ejercicios

import (
	"github.com/untref-ayp2/data-structures/queue"
	"github.com/untref-ayp2/data-structures/stack"
)

// Escribir una función que que reciba una cadena de caracteres
// y devuelva la cadena invertida. Analizar el orden.
func InvertirCadena(cadena string) string {
	invertida := ""
	cola := queue.New[string]()
	for i := len(cadena) - 1; i >= 0; i-- {
		cola.Enqueue(string(cadena[i]))
	}
	for !cola.IsEmpty() {
		aux, _ := cola.Dequeue()
		invertida += aux
	}
	return invertida
}

// Escribir una función que verifique si una palabra es palíndromo
// (es decir que una cadena es igual a su inversa. Por ejemplo:
//
//	las cadenas "1456541" y "145541" son palíndromos). Analizar el orden.
func Palindromo(cadena string) bool {
	pila := stack.New[string]()
	aux := ""
	for i := 0; i < len(cadena); i++ {
		pila.Push(string(cadena[i]))
	}
	for !pila.IsEmpty() {
		c, _ := pila.Pop()
		aux += c
	}
	return cadena == aux
}

// Escribir una función que evalúe si una cadena de paréntesis, corchetes y
// llaves está bien balanceada y devuelve `true` si está bien balanceada y
// `false` si está mal balanceada. Por ejemplo: `[()]{}{[()()]()}` debe
// devolver `true`, mientras que `[(])` debe devolver `false`. Analizar el orden.
func Balanceada(cadena string) bool {
	apertura := stack.New[string]()
	for i := 0; i < len(cadena); i++ {
		if string(cadena[i]) == "[" || string(cadena[i]) == "(" || string(cadena[i]) == "{" {
			apertura.Push(string(cadena[i]))
		} else if string(cadena[i]) == "]" || string(cadena[i]) == ")" || string(cadena[i]) == "}" {
			if apertura.IsEmpty() {
				return false
			}
			aux, _ := apertura.Pop()
			if aux != "[" && string(cadena[i]) == "]" || aux != "(" && string(cadena[i]) == ")" || aux != "{" && string(cadena[i]) == "}" {
				return false
			}
		}
	}
	return true
}

// Escribir una función, tal que, dadas dos colas, construya una cola con el resultado
// de poner una a continuación de la otra. Por ejemplo: si `q1 = [1,2,3]` (con 1 al
// frente y 3 al final) y `q2 = [5,7]`, el resultado es `[1,2,3,5,7]` (con 1 al
// frente y 7 al final). Analizar el orden.
func UnirColas[T any](q1, q2 *queue.Queue[T]) *queue.Queue[T] {
	nuevacola := queue.New[T]()
	for !q1.IsEmpty() || !q2.IsEmpty() {
		if !q1.IsEmpty() {
			aux, _ := q1.Dequeue()
			nuevacola.Enqueue(aux)
		} else {
			aux, _ := q2.Dequeue()
			nuevacola.Enqueue(aux)
		}
	}
	return nuevacola
}
