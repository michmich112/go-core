package maths

import (
	"github.com/gildas/go-errors"
)

/**
Matrix of dimension 2 type for input elements
*/
type Matrix []Vector32

/**
Creates a new identity matrix
*/
/*func IdentityMatrix(m int, n int) Matrix {
	mCount := 0
	for  i:=0 ; i<m; i++ {

	}
}*/

/**
Creates a new Identity matrix
*/
/*func (m *Matrix) Identity(mm int, n int) Matrix {

}
*/
/*
Verifies that two matrices are equal (deep equal)
*/
func (m *Matrix) Equals(b Matrix) bool {
	if len(*m) != len(b) {
		return false
	}

	equals := true
	for i := range *m {
		equals = equals && (*m)[i].Equals(b[i])
		if !equals {
			break // break if false to make the assertion faster
		}
	}
	return equals
}

/**
Add vector to matrix
*/
func (m *Matrix) AddVector(v Vector32) (mat Matrix, err error) {
	if len(*m) > 0 && len((*m)[0]) != len(v) {
		return *m, errors.New("Size Match error")
	}
	*m = append(*m, v)
	return *m, nil
}

/**
Transpose function with struct linking
*/
func (m *Matrix) Transpose() Matrix {
	*m = Transpose(*m)
	return *m
}

/**
Values range Struct
*/
type Range struct {
	Min float32
	Max float32
}

/**
Get The Range of values of the vectors in the matrix
*/
func (m *Matrix) GetValueRange() []Range {
	if len(*m) < 1 {
		return nil
	}
	t := Transpose(*m)
	l := len(t) // length of the vectors
	var r = make([]Range, l)
	for i := 0; i < l; i++ {
		min, max := MinMaxF32(t[i]...)
		r[i] = Range{min, max}
	}
	return r
}
