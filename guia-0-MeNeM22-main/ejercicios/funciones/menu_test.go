package funciones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElegirOpcionValida(t *testing.T) {
	// Probamos cada opción válida
	for i := 1; i <= 4; i++ {
		resultado := elegirOpcion(i)
		assert.Equal(t, i, resultado)
	}
}
func TestElegirOpcionInvalida(t *testing.T) {
	// Probamos una opción inválida
	resultado := elegirOpcion(5)
	assert.Equal(t, 0, resultado)
}
