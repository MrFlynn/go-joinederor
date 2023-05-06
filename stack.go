package joinederror

type stack[T any] struct {
	items []T
}

func newStack[T any]() *stack[T] {
	return &stack[T]{items: []T{}}
}

func (s *stack[T]) push(value T) int {
	s.items = append(s.items, value)
	return len(s.items)
}

func (s *stack[T]) pop() (T, int) {
	l := len(s.items)

	if l == 0 {
		var nothing T
		return nothing, 0
	}

	last := s.items[l-1]
	s.items = s.items[:l-1]

	return last, l - 1
}

func (s *stack[T]) empty() bool {
	return len(s.items) == 0
}
