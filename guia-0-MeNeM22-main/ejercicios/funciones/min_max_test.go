package funciones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncontrarMinimoMaximo(t *testing.T) {
	lista := []int{4, 7, 2, 9, 1}

	minimo, maximo := EncontrarMinimoMaximo(lista)
	assert.Equal(t, 1, minimo)
	assert.Equal(t, 9, maximo)
}
