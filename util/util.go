package util

// MinInt returns the smaller of x or y.
func MinInt(a ,b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt returns the larger of x or y.
func MaxInt(a ,b int) int {
	if a > b {
		return a
	}
	return b
}