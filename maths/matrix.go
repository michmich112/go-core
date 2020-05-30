package maths

import (
	"github.com/gildas/go-errors"
)

/**
Matrix of dimension 2 type for input elements
*/
type Matrix []Vector

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
func (m *Matrix) AddVector(v Vector) (mat Matrix, err error) {
	if len((*m)[0]) != len(v) {
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
