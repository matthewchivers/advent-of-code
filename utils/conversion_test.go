package utils

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		err      error
	}{
		{"1", 1, nil},
		{"-1", -1, nil},
		{"0", 0, nil},
		{"-0", 0, nil},
		{"-123", -123, nil},
		{"123", 123, nil},
		{"-1234567890", -1234567890, nil},
		{"1234567890", 1234567890, nil},
		{"", 0, strconv.ErrSyntax},
		{"-1.0", 0, strconv.ErrSyntax},
		{"1.0", 0, strconv.ErrSyntax},
	}

	for _, test := range tests {
		result, err := StringToInt(test.input)
		if test.err != nil {
			assert.Error(t, test.err, err)
		}
		assert.Equal(t, test.expected, result, "StringToInt() should convert the string to the correct int value")
	}
}
