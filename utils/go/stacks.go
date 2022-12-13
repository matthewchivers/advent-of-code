package adventofcode

// RuneStack is a stack of runes
type RuneStack struct {
	items []rune
}

// Push adds an item to the stack
func (s *RuneStack) Push(r rune) {
	// First in, last out
	s.items = append(s.items, r)
}

// Pop removes an item from the stack
func (s *RuneStack) Pop() rune {
	if len(s.items) == 0 {
		return ' '
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// LIFOPush (pushes to bottom of the stack)
func (s *RuneStack) LIFOPush(r rune) {
	s.items = append([]rune{r}, s.items...)
}

// Len returns the length of the stack
func (s *RuneStack) Len() int {
	return len(s.items)
}

// GetStack returns the stack as a string
func (s *RuneStack) GetStack() string {
	output := ""
	for _, item := range s.items {
		output += string(item) + ", "
	}
	return output
}

// Peek returns the top item on the stack without removing it from the stack
func (s *RuneStack) Peek() rune {
	if len(s.items) == 0 {
		return ' '
	}
	return s.items[len(s.items)-1]
}
