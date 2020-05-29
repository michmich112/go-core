package maths

/**
Vector of int
*/
type Vector []float64

/**
Verifies that two vectors are equal (deep equal)
*/
func (a *Vector) Equals(b Vector) bool {
	if len(*a) != len(b) {
		return false
	}

	equals := true
	for i := range *a {
		equals = equals && ((*a)[i] == b[i])
		if !equals {
			break
		} // break if false to make the assertion faster
	}
	return equals
}

/**
Matrix of dimension 2 type for input elements
*/
type Matrix []Vector

/*
Verifies that two matrices are equal (deep equal)
*/
func (a *Matrix) Equals(b Matrix) bool {
	if len(*a) != len(b) {
		return false
	}

	equals := true
	for i := range *a {
		equals = equals && (*a)[i].Equals(b[i])
		if !equals {
			break // break if false to make the assertion faster
		}
	}
	return equals
}
