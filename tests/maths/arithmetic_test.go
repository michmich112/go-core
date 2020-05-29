package maths

import (
	"core/maths"
	"fmt"
	"testing"
)

/**
Transpose Tests
*/
func TestTranspose(t *testing.T) {
	/**
	    | 1 2 3 4 |    | 1 5 9 |
	  m | 5 6 7 8 | -> | 2 6 0 |
	    | 9 0 1 2 |    | 3 7 1 |
	                   | 4 8 2 |
	*/
	mInitial := maths.Matrix{
		[]float64{1, 2, 3, 4},
		[]float64{5, 6, 7, 8},
		[]float64{9, 0, 1, 2}}

	mExpected := maths.Matrix{
		[]float64{1, 5, 9},
		[]float64{2, 6, 0},
		[]float64{3, 7, 1},
		[]float64{4, 8, 2}}

	mInitial.Equals(mExpected)
	transpose := maths.Transpose(mInitial)
	if !transpose.Equals(mExpected) {
		fmt.Print("Expected:", mExpected, "Received:", mInitial)
		t.Errorf("Received did not match expected")
	}
}
