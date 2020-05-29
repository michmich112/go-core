package core

/**
Similar to the python `in range(number)`
*/
func InRange(bound *int) bool {
	*bound -= 1
	return *bound >= 0
}

/**
In range ascending
]variable,bound]
*/
func InRangeAsc(variable *int, bound int) bool {
	*variable++
	return *variable <= bound
}

/**
Array Map function
*/
func Map(array []interface{}, f func(i interface{}) interface{}) []interface{} {
	for i, item := range array {
		array[i] = f(item)
	}
	return array
}

///**
//Applies a function through items of a vector
// */
//func ForEach(array maths.Vector, f func(i float64)) {
//	for _, item := range array {
//		f(item)
//	}
//}
