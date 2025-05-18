package guia4

import (
	"github.com/untref-ayp2/data-structures/set"
	"github.com/untref-ayp2/data-structures/types"
)

func Union[T types.Ordered](s1, s2 *set.SetList[T]) *set.SetList[T] {
	newSet := set.NewSetList[T]()
	for _, aux1 := range s1.Values() {
		newSet.Add(aux1)
	}
	for _, aux2 := range s2.Values() {
		if !newSet.Contains(aux2) {
			newSet.Add(aux2)
		}
	}
	return newSet
}

func Intersection[T types.Ordered](s1, s2 *set.SetList[T]) *set.SetList[T] {
	newSet := set.NewSetList[T]()
	for _, aux1 := range s1.Values() {
		if s2.Contains(aux1) {
			newSet.Add(aux1)
		}
	}
	return newSet
}

func Difference[T types.Ordered](s1, s2 *set.SetList[T]) *set.SetList[T] {
	newSet := set.NewSetList[T]()
	for _, aux1 := range s1.Values() {
		if !s2.Contains(aux1) {
			newSet.Add(aux1)
		}
	}
	return newSet
}

func Subset[T types.Ordered](s1, s2 *set.SetList[T]) bool {
	for _, aux1 := range s2.Values() {
		if !s1.Contains(aux1) {
			return false
		}
	}
	return true
}

func Equal[T types.Ordered](s1, s2 *set.SetList[T]) bool {
	for _, aux1 := range s1.Values() {
		if !s2.Contains(aux1) {
			return false
		}
	}
	return true
}

func SymmetricDifference[T types.Ordered](s1, s2 *set.SetList[T]) *set.SetList[T] {
	newSet := set.NewSetList[T]()
	for _, aux1 := range s1.Values() {
		if !s2.Contains(aux1) {
			newSet.Add(aux1)
		}
	}
	for _, aux2 := range s2.Values() {
		if !s1.Contains(aux2) {
			newSet.Add(aux2)
		}
	}
	return newSet
}

func EliminarRepetidos[T types.Ordered](array []T) []T {
	newArray := []T{}
	existe := map[T]bool{}

	for _, v := range array {
		// Solo añadimos el elemento si no ha sido agregado anteriormente
		if !existe[v] {
			newArray = append(newArray, v)
			existe[v] = true
		}
	}
	return newArray
}

func InterseccionMultiple[T types.Ordered](sets ...*set.SetList[T]) *set.SetList[T] {
	// Implementar
	return nil
}
