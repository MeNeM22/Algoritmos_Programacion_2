package funciones

import "fmt"

// Pre: El parámetro coeficientes no debe ser nulo.
// Post: La cadena devuelta representa el polinomio formado por los coeficientes.
func ImprimirPolinomio(coeficientes []float64) string {
	polinomio := ""
	if len(coeficientes) != 0 {
		for i := 0; i < len(coeficientes); i++ {
			if i == 0 {
				polinomio += fmt.Sprintf("%g ", coeficientes[0]) + "+"
			}
			if i == 1 {
				polinomio += fmt.Sprintf(" %g x ", coeficientes[1])
			}
			if i > 1 {
				polinomio += fmt.Sprintf("+ %g x**%d", coeficientes[i], i)
			}
		}
	}
	return polinomio
}
