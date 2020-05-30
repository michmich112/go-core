package maths

/**
Vector of int
*/
type Vector []float64

/**
Verifies that two vectors are equal (deep equal)
*/
func (v *Vector) Equals(b Vector) bool {
	if len(*v) != len(b) {
		return false
	}

	equals := true
	for i := range *v {
		equals = equals && ((*v)[i] == b[i])
		if !equals {
			break
		} // break if false to make the assertion faster
	}
	return equals
}

/**
Push new value into the vector
*/
func (v *Vector) Push(val float64) {
	*v = append(*v, val)
}
