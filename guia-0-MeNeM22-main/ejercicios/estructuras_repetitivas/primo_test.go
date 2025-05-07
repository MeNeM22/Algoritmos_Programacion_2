package estructurasrepetitivas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimo(t *testing.T) {
	resultado := EsPrimo(5)
	assert.Equal(t, true, resultado)
}

func TestNoPrimo(t *testing.T) {
	resultado := EsPrimo(4)
	assert.Equal(t, false, resultado)
}
