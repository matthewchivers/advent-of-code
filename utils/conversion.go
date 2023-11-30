package utils

import (
	"strconv"
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
