package maths

import (
	"go-core"
	"math"
)

/**
x^2
*/
func Square(x float64) float64 {
	return math.Pow(x, 2)
}

/**
x^3
*/
func Cube(x float64) float64 {
	return math.Pow(x, 3)
}

/**
Mean of a Vector64 (array of float64's)
*/
func Mean(array Vector32) float32 {
	l := float32(len(array))
	var s float32
	for i := range array {
		s += array[i] / l
	}
	return s
}

/**
Translates matrix to new dimensions
*/
func Transpose(matrix Matrix) Matrix {
	m := len(matrix)
	n := len(matrix[0])

	i := 0

	newMatrix := Matrix{}
	for core.InRangeAsc(&i, n) {
		vector := Vector32{}
		j := 0
		for core.InRangeAsc(&j, m) {
			vector = append(vector, matrix[j-1][i-1])
		}
		newMatrix = append(newMatrix, vector)
	}
	return newMatrix
}
