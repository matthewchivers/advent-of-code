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

func TestStringToIntArray(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []int
		wantErr bool
	}{
		{
			"empty string",
			"",
			[]int{},
			false,
		},
		{
			"single value",
			"1",
			[]int{1},
			false,
		},
		{
			"multiple values",
			"1 2 3 4 5",
			[]int{1, 2, 3, 4, 5},
			false,
		},
		{
			"multiple values with negative",
			"1 -2 3 -4 5",
			[]int{1, -2, 3, -4, 5},
			false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToIntArray(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToIntArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got, "StringToIntArray() should convert the string to the correct int array")
		})
	}
}
