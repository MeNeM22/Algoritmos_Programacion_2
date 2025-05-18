package guia4

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/untref-ayp2/data-structures/set"
)

func TestUnion(t *testing.T) {
	set1 := set.NewSetList(1, 2)
	set2 := set.NewSetList(1, 3)
	result := Union(set1, set2)
	require.NotNil(t, result)
	assert.Equal(t, 3, result.Size(), "La union de set1 y set2 deberia tener 3 elementos")
	esperado := []int{3, 1, 2}
	values := result.Values()
	for _, v := range esperado {
		assert.Contains(t, values, v, "La union de set1 y set2 deberia dar [1,2,3]")
	}
	set1 = set.NewSetList[int]()
	set2 = set.NewSetList(1, 3)
	result = Union(set1, set2)
	require.NotNil(t, result)
	assert.Equal(t, 2, result.Size(), "La union de set1 y set2 deberia tener 2 elementos")
	esperado = []int{3, 1}
	values = result.Values()
	for _, v := range esperado {
		assert.Contains(t, values, v, "La union de set1 y set2 deberia dar [1,3]")
	}

}

func TestIntersection(t *testing.T) {
	set1 := set.NewSetList(1, 2, 5)
	set2 := set.NewSetList(1, 3, 5)
	result := Intersection(set1, set2)
	require.NotNil(t, result)
	assert.Equal(t, 2, result.Size(), "La intersección de set1 y set2 deberia tener 2 elementos")
	esperado := []int{1, 5}
	values := result.Values()
	for _, v := range esperado {
		assert.Contains(t, values, v, "La intersección de set1 y set2 deberia dar [1,5]")
	}
	set1 = set.NewSetList[int]()
	set2 = set.NewSetList(1, 3, 5)
	result = Intersection(set1, set2)
	require.NotNil(t, result)
	assert.Equal(t, 0, result.Size(), "La intersección de set1 y set2 deberia tener 0 elementos")
}

func TestDifference(t *testing.T) {
	set1 := set.NewSetList(1, 2, 4, 7)
	set2 := set.NewSetList(1, 3, 7)
	result := Difference(set1, set2)
	require.NotNil(t, result)
	assert.Equal(t, 2, result.Size(), "La differencia de set1 y set2 deberia tener 2 elementos")
	esperado := []int{2, 4}
	values := result.Values()
	for _, v := range esperado {
		assert.Contains(t, values, v, "La differencia de set1 y set2 deberia dar [2,4]")
	}
	set1 = set.NewSetList(1, 2, 4, 7)
	set2 = set.NewSetList[int]()
	result = Difference(set1, set2)
	require.NotNil(t, result)
	assert.Equal(t, 4, result.Size(), "La differencia de set1 y set2 deberia tener 4 elementos")
	esperado = []int{1, 2, 4, 7}
	values = result.Values()
	for _, v := range esperado {
		assert.Contains(t, values, v, "La differencia de set1 y set2 deberia dar [1, 2, 4, 7]")
	}
}

func TestSubset(t *testing.T) {
	set1 := set.NewSetList(1, 2)
	set2 := set.NewSetList(1)
	assert.True(t, Subset(set1, set2), "set2 deberia ser un subconjunto de set1")
}
func TestNoSubset(t *testing.T) {
	set1 := set.NewSetList(1, 2)
	set2 := set.NewSetList(5)
	assert.False(t, Subset(set1, set2), "set2 no deberia ser un subconjunto de set1")
}

func TestEqual(t *testing.T) {
	set1 := set.NewSetList(1, 2)
	set2 := set.NewSetList(2, 1)
	assert.True(t, Equal(set1, set2), "Los conjuntos deberian ser iguales")
}

func TestNotEqual(t *testing.T) {
	set1 := set.NewSetList(1, 2)
	set2 := set.NewSetList(2, 3)
	assert.False(t, Equal(set1, set2), "Los conjuntos deberian ser distintos")
}

func TestSymmetricDifference(t *testing.T) {
	set1 := set.NewSetList(1, 2, 3, 4, 7, 11)
	set2 := set.NewSetList(3, 4, 11, 12)
	result := SymmetricDifference(set1, set2)
	require.NotNil(t, result)
	assert.Equal(t, 4, result.Size(), "La diferencia simetrica de set1 y set2 deberia tener 4 elementos")
	esperado := []int{1, 2, 7, 12}
	values := result.Values()
	for _, v := range esperado {
		assert.Contains(t, values, v, "La diferencia simetrica de set1 y set2 deberia dar [1,2,7]")
	}
	set1 = set.NewSetList(1, 2, 3, 4, 7, 11)
	set2 = set.NewSetList[int]()
	result = SymmetricDifference(set1, set2)
	require.NotNil(t, result)
	assert.Equal(t, 6, result.Size(), "La diferencia simetrica de set1 y set2 deberia tener 6 elementos")
	esperado = []int{1, 2, 3, 4, 7, 11}
	values = result.Values()
	for _, v := range esperado {
		assert.Contains(t, values, v, "La diferencia simetrica de set1 y set2 deberia dar [1, 2, 3, 4, 7, 11]")
	}
}

func TestEliminarRepetidos(t *testing.T) {
	array := []int{1, 2, 7, 2, 1, 3}
	result := EliminarRepetidos(array)
	require.NotNil(t, result)
	assert.Equal(t, 4, len(result), "El array deberia tener 3 elementos")
	esperado := []int{1, 2, 7, 3}
	for _, v := range esperado {
		assert.Contains(t, result, v, "El array deberia dar [1,2,7,3]")
	}
}

func TestInterseccionMultiple(t *testing.T) {
	set1 := set.NewSetList(1, 2, 3)
	set2 := set.NewSetList(2, 3)
	set3 := set.NewSetList(2)

	result := InterseccionMultiple(set1, set2, set3)
	require.NotNil(t, result)
	assert.Equal(t, 1, result.Size(), "La interseccion entre set1, set2 y set3 deberia tener 1 elemento")
	esperado := []int{2}
	values := result.Values()
	for _, v := range esperado {
		assert.Contains(t, values, v, "La interseccion entre set1, set2 y set3 deberia dar [2]")
	}
	set1 = set.NewSetList(7, 11)
	set2 = set.NewSetList[int]()
	set3 = set.NewSetList[int]()
	result = InterseccionMultiple(set1, set2, set3)
	require.NotNil(t, result)
	assert.Equal(t, 0, result.Size(), "La interseccion entre set1, set2 y set3 deberia tener 0 elementos")
	set1 = set.NewSetList(7, 11)
	result = InterseccionMultiple(set1)
	require.NotNil(t, result)
	assert.Equal(t, 0, result.Size(), "La interseccion con un solo conjunto deberia tener 0 elementos")
}
