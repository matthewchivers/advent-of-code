package utils

// MaxInt returns the larger of two integers
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinInt returns the smaller of two integers
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
