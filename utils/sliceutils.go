package utils

import "fmt"

// SlicePop removes the last element from an *[]T and returns it
func SlicePop[T any](s *[]T) (T, error) {
	if len(*s) == 0 {
		var zeroValue T // zeroValue is the zero value of type T
		return zeroValue, fmt.Errorf("SlicePop() cannot pop from an empty slice")
	}
	retVal := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return retVal, nil
}

// SliceContainsRune returns true if a rune is in a slice of runes
func SliceContainsRune(s []rune, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}

// SliceContainsString returns true if a string is in a slice of strings
func SliceContainsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
