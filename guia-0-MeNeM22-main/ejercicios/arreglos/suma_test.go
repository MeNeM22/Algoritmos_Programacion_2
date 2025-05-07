package arreglos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuma(t *testing.T) {
	lista := []int{4, 7, 2, 9, 1}

	resultado := Suma(lista)
	assert.Equal(t, 23, resultado)
}
