package arreglos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumaProductoEscalar(t *testing.T) {
	v1 := []int{1, 2, 3}
	v2 := []int{4, 5, 6}

	suma, producto := SumaProductoEscalar(v1, v2)
	assert.Equal(t, 21, suma)
	assert.Equal(t, 32, producto)
}
