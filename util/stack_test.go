package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	stackInt Stack[int]
)

func init() {
	// Initialize stackInt with values
	for _, v := range []int{1, 2, 3} {
		stackInt.Push(v)
	}
}

func TestStackPush(t *testing.T) {
	stack := stackInt

	// Check if the stack contains the expected elements
	assert.Equal(t, []int{1, 2, 3}, stack.items, "Push() should add elements to the stack")

	// Push another element onto the stack
	stack.Push(4)

	// Check if the stack contains the updated elements
	assert.Equal(t, []int{1, 2, 3, 4}, stack.items, "Push() should add elements to the stack")
}

func TestStackPop(t *testing.T) {
	stack := stackInt

	// Pop an element from the stack
	popped := stack.Pop()

	// Check if the popped element is the expected value
	assert.Equal(t, 3, popped, "Pop() should return the top element from the stack")

	// Check if the stack contains the updated elements
	assert.Equal(t, []int{1, 2}, stack.items, "Pop() should remove the top element from the stack")
}

func TestLen(t *testing.T) {
	stack := stackInt

	// Check if the stack has the expected length
	assert.Equal(t, 3, stack.Len(), "Len() should return the length of the stack")

	// Pop an element from the stack
	stack.Pop()

	// Check if the stack has the expected length
	assert.Equal(t, 2, stack.Len(), "Len() should return the length of the stack")
}

func TestGetStackString(t *testing.T) {
	stack := stackInt

	// Check if the stack has the expected string representation
	assert.Equal(t, "[ 1, 2, 3 ]", stack.GetStackString(), "GetStackString() should return the stack as a string")
}

func TestPeek(t *testing.T) {
	stack := stackInt

	// Check if the peeked element is the expected value
	assert.Equal(t, 3, stack.Peek(), "Peek() should return the top element from the stack")

	// Pop an element from the stack
	stack.Pop()

	// Check if the peeked element is the expected value
	assert.Equal(t, 2, stack.Peek(), "Peek() should return the top element from the stack")
}

func TestReverse(t *testing.T) {
	stack := stackInt

	// Reverse the stack
	stack.Reverse()

	// Check if the stack has the expected elements
	assert.Equal(t, []int{3, 2, 1}, stack.items, "Reverse() should reverse the order of the stack")
}
