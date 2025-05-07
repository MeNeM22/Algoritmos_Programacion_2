package estructurasrepetitivas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	resultado := Factorial(5)
	assert.Equal(t, 120, resultado)
}
