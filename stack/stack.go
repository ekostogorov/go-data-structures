package stack

// Stack represents stack structure
type Stack struct {
	size    int
	storage []string
}

// New constructs stack
func New() *Stack {
	return &Stack{}
}

// Push adds element to stack
func (s *Stack) Push(element string) {
	s.storage = append(s.storage, element)
	s.size++
}

// Pop removes element from stack
func (s *Stack) Pop() string {
	if len(s.storage) == 0 {
		panic("State is empty")
	}

	element := s.storage[s.size-1]
	s.storage = s.storage[:s.size-1]
	s.size--

	return element
}

// Size returns length of stack
func (s *Stack) Size() int {
	return s.size
}
