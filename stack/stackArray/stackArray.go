package stackArray

import "errors"

const stackSize = 5

type StackArray[T any] struct {
	stack [stackSize]T
	top   int
}

// push
func (s *StackArray[T]) Push(x T) error {
	// check for stack overflow
	if s.IsFull() {
		return errors.New("stack is full")
	}

	s.top++
	s.stack[s.top] = x
	return nil
}

// pop top node in the stack
func (s *StackArray[T]) Pop() (T, error) {
	// check for stack underflow
	if s.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}

	// save top node
	topNode := s.stack[s.top]
	s.top--
	return topNode, nil
}

// peek top node in the stack
func (s *StackArray[T]) Peek() (T, error) {
	// check for stack underflow
	if s.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}
	return s.stack[s.top], nil
}

// isEmpty
func (s *StackArray[T]) IsEmpty() bool {
	return s.top == -1
}

// isFull
func (s *StackArray[T]) IsFull() bool {
	return s.top >= stackSize
}
