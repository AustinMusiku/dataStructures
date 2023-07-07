package stackArray

import "errors"

const stackSize = 5

type StackArray[T any] struct {
	stack [stackSize]T
	top   int
}

// NewStack initialize new Stack
func NewStack[T any]() *StackArray[T] {
	stack := new(StackArray[T])
	stack.top = -1
	return stack
}

// Push add node to the top of the stack
func (s *StackArray[T]) Push(x T) error {
	// check for stack overflow
	if s.IsFull() {
		return errors.New("stack is full")
	}
	s.top++
	s.stack[s.top] = x
	return nil
}

// Pop remove top node in the stack
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

// Peek peek top node in the stack
func (s *StackArray[T]) Peek() (T, error) {
	// check for stack underflow
	if s.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}
	return s.stack[s.top], nil
}

// IsEmpty check if stack is empty
func (s *StackArray[T]) IsEmpty() bool {
	return s.top == -1
}

// IsFull check if stack is full
func (s *StackArray[T]) IsFull() bool {
	return s.top >= stackSize
}

// Print log nodes in the stack
func (s *StackArray[T]) Print() {
	// print each node with its index
	for i := s.top; i >= 0; i-- {
		println(i, s.stack[i])
	}
}
