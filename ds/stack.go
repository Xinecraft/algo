package ds

import "fmt"

type Stack[T comparable] struct {
	stack  LinkedList[T]
	length int
}

func (s *Stack[T]) Push(item T) {
	s.length++
	s.stack.Prepend(item)
}

func (s *Stack[T]) Pop() T {
	if s.length <= 0 {
		return *new(T)
	}

	s.length--
	return s.stack.RemoveFirst()
}

func (s *Stack[T]) Peek(item T) T {
	node := s.stack.head

	if node == nil {
		return *new(T)
	}

	return node.value
}

func (s *Stack[T]) Size(item T) int {
	return s.length
}

func (s Stack[T]) String() string {
	str := fmt.Sprint(s.stack)
	return str
}
