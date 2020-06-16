package utils

// FilterInt function
func FilterInt(arr []int, cb func(int) bool) []int {
	rv := arr[:0]
	for _, v := range arr {
		if cb(v) {
			rv = append(rv, v)
		}
	}
	return rv
}

// FilterFloat64 function
func FilterFloat64(arr []float64, cb func(float64) bool) []float64 {
	rv := arr[:0]
	for _, v := range arr {
		if cb(v) {
			rv = append(rv, v)
		}
	}
	return rv
}

// FilterString function
func FilterString(arr []string, cb func(string) bool) []string {
	rv := arr[:0]
	for _, v := range arr {
		if cb(v) {
			rv = append(rv, v)
		}
	}
	return rv
}
