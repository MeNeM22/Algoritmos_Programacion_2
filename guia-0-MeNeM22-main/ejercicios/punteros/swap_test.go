package punteros

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwap(t *testing.T) {
	a, b := 1, 2

	Swap(&a, &b)
	assert.Equal(t, 2, a)
	assert.Equal(t, 1, b)
}
