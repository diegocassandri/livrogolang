package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Sem generics
func MapValues(values []int, mapFunc func(int) int) []int {
	var newValues []int
	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

// comGenerics
func MapValuesGeneric[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func main() {
	result := MapValues([]int{1, 2, 3}, func(n int) int {
		return n * 3
	})

	fmt.Printf("result %+v", result)

	result2 := MapValuesGeneric([]float64{1, .0, 2.9, 3.5}, func(n float64) float64 {
		return n * 3
	})

	fmt.Printf("result %+v", result2)
}
