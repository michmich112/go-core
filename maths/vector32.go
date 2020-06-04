package maths

/**
Vector32 of int
*/
type Vector32 []float32

/**
Verifies that two vectors are equal (deep equal)
*/
func (v Vector32) Equals(b Vector32) bool {
	if len(v) != len(b) {
		return false
	}

	equals := true
	for i := range v {
		equals = equals && ((v)[i] == b[i])
		if !equals {
			break
		} // break if false to make the assertion faster
	}
	return equals
}

/**
Push new value into the vector
*/
func (v *Vector32) Push(val float32) {
	*v = append(*v, val)
}
