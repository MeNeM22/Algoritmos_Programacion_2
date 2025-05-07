package estructurasrepetitivas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProducto(t *testing.T) {
	resultado := Producto(3, 4)
	assert.Equal(t, 12, resultado)
}
