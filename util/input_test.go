package util

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileAsLines(t *testing.T) {
	tests := []struct {
		name     string
		testData string
	}{
		{
			name:     "one line",
			testData: "Hello, World!",
		},
		{
			name:     "multiple lines",
			testData: "Hello, World 1!\nHello, World 2!\n Hello, World 3",
		},
		{
			name:     "empty file",
			testData: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file for testing
			tempFile, err := os.CreateTemp("", "test")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tempFile.Name())

			// Write test data to the temporary file
			err = os.WriteFile(tempFile.Name(), []byte(tt.testData), 0644)
			if err != nil {
				t.Fatal(err)
			}

			// Call the function being tested
			result := ReadFileAsLines(tempFile.Name())

			// Check if the result is nil (empty file)
			if tt.testData == "" {
				assert.Nil(t, result)
				return
			}

			// Create the expected result
			expected := strings.Split(tt.testData, "\n")

			// Compare the two slices
			assert.Equal(t, expected, result)
		})
	}
}

func TestReadFileAsBytes(t *testing.T) {
	tests := []struct {
		name     string
		testData string
	}{
		{
			name:     "one line",
			testData: "Hello, World!",
		},
		{
			name:     "multiple lines",
			testData: "Hello, World 1!\nHello, World 2!\n Hello, World 3",
		},
		{
			name:     "empty file",
			testData: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file for testing
			tempFile, err := os.CreateTemp("", "test")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tempFile.Name())

			// Write test data to the temporary file
			err = os.WriteFile(tempFile.Name(), []byte(tt.testData), 0644)
			if err != nil {
				t.Fatal(err)
			}

			// Call the function being tested
			result := ReadFileAsBytes(tempFile.Name())

			// Compare the two slices
			assert.Equal(t, []byte(tt.testData), result)
		})
	}
}

func TestReadFileAsByteChunks(t *testing.T) {
	tests := []struct {
		name     string
		fileData string
		want     [][]byte
	}{
		{
			name:     "one chunk",
			fileData: "Hello, World!",
			want: [][]byte{
				[]byte("Hello, World!"),
			},
		},
		{
			name:     "multiple chunks",
			fileData: "Hello, World 1!\n\nHello, World 2!\n\nHello, World 3!",
			want: [][]byte{
				[]byte("Hello, World 1!"),
				[]byte("Hello, World 2!"),
				[]byte("Hello, World 3!"),
			},
		},
		{
			name:     "multiline chunks",
			fileData: "Hello, World 1!\nHello, World 2!\n\nHello, World 3\nHello, World 4!\n",
			want: [][]byte{
				[]byte("Hello, World 1!\nHello, World 2!"),
				[]byte("Hello, World 3\nHello, World 4!"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file for testing
			tempFile, err := os.CreateTemp("", "test")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tempFile.Name())

			// Write test data to the temporary file
			err = os.WriteFile(tempFile.Name(), []byte(tt.fileData), 0644)
			if err != nil {
				t.Fatal(err)
			}

			// Call the function being tested
			actual := ReadFileAsByteChunks(tempFile.Name())

			// Compare the two slices
			assert.Equal(t, tt.want, actual, "ReadFileAsByteChunks() actual = %v, wanted %v", actual, tt.want)
		})
	}
}
