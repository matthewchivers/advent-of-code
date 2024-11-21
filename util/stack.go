package util

import "fmt"

// Stack is a generic stack implementation
type Stack[T any] struct {
	items []T
}

// Push adds an item to the stack
func (s *Stack[T]) Push(r T) {
	s.items = append(s.items, r)
}

// Pop removes an item from the stack
func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zeroValue T
		return zeroValue
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Len returns the length of the stack
func (s *Stack[T]) Len() int {
	return len(s.items)
}

// GetStack returns the stack as a string
func (s *Stack[T]) GetStackString() string {
	if len(s.items) == 0 {
		return "[ ]"
	}
	output := "[ "
	for _, item := range s.items {
		if output != "[ " {
			output += ", "
		}
		output += fmt.Sprintf("%v", item)
	}
	output += " ]"
	return output
}

// Peek returns the top item on the stack without removing it from the stack
func (s *Stack[T]) Peek() T {
	if len(s.items) == 0 {
		var zeroValue T
		return zeroValue
	}
	return s.items[len(s.items)-1]
}

// Reverse reverses the order of the stack
func (s *Stack[T]) Reverse() {
	for i := 0; i < len(s.items)/2; i++ {
		j := len(s.items) - i - 1
		s.items[i], s.items[j] = s.items[j], s.items[i]
	}
}
