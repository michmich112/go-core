package maths

/**
Vector64 of int
*/
type Vector64 []float64

/**
Verifies that two vectors are equal (deep equal)
*/
func (v *Vector64) Equals(b Vector64) bool {
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
func (v *Vector64) Push(val float64) {
	*v = append(*v, val)
}
