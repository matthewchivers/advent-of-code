package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlicePop(t *testing.T) {
	testCases := []struct {
		name      string
		slice     interface{}
		expected  interface{}
		expectErr error
	}{
		{
			name:      "Empty string slice",
			slice:     []string{},
			expected:  "",
			expectErr: fmt.Errorf("SlicePop() cannot pop from an empty slice"),
		},
		{
			name:      "Non-empty string slice",
			slice:     []string{"a", "b", "c"},
			expected:  "c",
			expectErr: nil,
		},
		{
			name:      "Empty int slice",
			slice:     []int{},
			expected:  0,
			expectErr: fmt.Errorf("SlicePop() cannot pop from an empty slice"),
		},
		{
			name:      "Non-empty int slice",
			slice:     []int{1, 2, 3},
			expected:  3,
			expectErr: nil,
		},
		{
			name:      "Single element string slice",
			slice:     []string{"a"},
			expected:  "a",
			expectErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			switch s := tc.slice.(type) {
			case []int:
				result, err := SlicePop(&s)
				if tc.expectErr != nil {
					assert.Equal(t, tc.expectErr, err, "SlicePop() should return the expected error")
					break
				} else {
					assert.Nil(t, err, "SlicePop() should not return an error")
				}
				assert.Equal(t, tc.expected, result, "SlicePop() should return the expected value")
			case []string:
				result, err := SlicePop(&s)
				if tc.expectErr != nil {
					assert.Equal(t, tc.expectErr, err, "SlicePop() should return the expected error")
					break
				} else {
					assert.Nil(t, err, "SlicePop() should not return an error")
				}
				assert.Equal(t, tc.expected, result, "SlicePop() should return the expected value")
			}
		})
	}
}

func TestSliceContainsRune(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []rune
		r        rune
		expected bool
	}{
		{
			name:     "Empty rune slice - no match",
			slice:    []rune{},
			r:        'a',
			expected: false,
		},
		{
			name:     "Non-empty rune slice - match",
			slice:    []rune{'a', 'b', 'c'},
			r:        'b',
			expected: true,
		},
		{
			name:     "Non-empty rune slice - no match",
			slice:    []rune{'a', 'b', 'c'},
			r:        'd',
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SliceContainsRune(tc.slice, tc.r)
			assert.Equal(t, tc.expected, result, "SliceContainsRune() should return the expected value")
		})
	}
}

func TestSliceContainsString(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []string
		str      string
		expected bool
	}{
		{
			name:     "Empty string slice - no match",
			slice:    []string{},
			str:      "a",
			expected: false,
		},
		{
			name:     "Non-empty string slice - match",
			slice:    []string{"a", "b", "c"},
			str:      "b",
			expected: true,
		},
		{
			name:     "Non-empty string slice - no match",
			slice:    []string{"a", "b", "c"},
			str:      "d",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SliceContainsString(tc.slice, tc.str)
			assert.Equal(t, tc.expected, result, "SliceContainsString() should return the expected value")
		})
	}
}
