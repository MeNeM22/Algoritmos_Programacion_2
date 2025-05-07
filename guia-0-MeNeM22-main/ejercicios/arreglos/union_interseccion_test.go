package arreglos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionInterseccion(t *testing.T) {
	lista1 := []int{1, 2, 3, 4, 5}
	lista2 := []int{3, 4, 5, 6, 7}

	union, interseccion := UnionInterseccion(lista1, lista2)
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7}, union)
	assert.ElementsMatch(t, []int{3, 4, 5}, interseccion)
}
