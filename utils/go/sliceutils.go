package adventofcode

// SlicePopInt removes the last element from an *[]int and returns it
func SlicePopInt(s *[]int) int {
	backVal := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return backVal
}
