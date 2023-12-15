package utils

import (
	"strconv"
	"strings"
)

// StringToInt converts a string to an int
// wrapper for strconv.Atoi for standardization / readability
func StringToInt(line string) (int, error) {
	val, err := strconv.Atoi(line)
	if err != nil {
		return 0, err
	}
	return val, nil
}

// StringToIntArray converts a string to an array of ints
// wrapper for strconv.Atoi for standardization / readability
func StringToIntArray(line string) ([]int, error) {
	var intArray []int
	fields := strings.Fields(line)
	for _, field := range fields {
		val, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		intArray = append(intArray, val)
	}
	return intArray, nil
}

// BoolToInt converts a bool to an int
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
