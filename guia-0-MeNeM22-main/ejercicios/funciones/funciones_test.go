package funciones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImprimirPolinomio(t *testing.T) {
	coeficientes := []float64{3.5, 2.5, 1.5, 2.6}
	resultadoEsperado := "3.5 + 2.5 x + 1.5 x**2 + 2.6 x**3"
	resultado := ImprimirPolinomio(coeficientes)
	assert.Equal(t, resultadoEsperado, resultado)
}
func TestImprimirPolinomioVacio(t *testing.T) {
	coeficientes := []float64{}
	resultadoEsperado := "El array está vacío"
	resultado := ImprimirPolinomio(coeficientes)
	assert.Equal(t, resultadoEsperado, resultado)
}
