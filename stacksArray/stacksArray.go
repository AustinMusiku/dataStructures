package stacksArray

import "errors"

const stackSize = 5

type StackArray struct{
	stack [stackSize]int
	top int
}

// push
func (s *StackArray) Push(x int) error {
	// check for stack overflow
	if s.IsFull(){
		return errors.New("stack is full")
	}

	s.top++
	s.stack[s.top] = x
	return nil
}

// pop top node in the stack
func (s *StackArray) Pop() (int, error){
	// check for stack underflow
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	// save top node
	top := s.stack[s.top]
	s.stack[s.top] = 0
	s.top--

	return top, nil
}

// peek top node in the stack
func (s *StackArray) Peek() (int, error){
	// check for stack underflow
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	return s.stack[s.top], nil
}

// isEmpty
func (s *StackArray) IsEmpty() bool{
	return s.top == -1
}

// isFull
func (s *StackArray) IsFull() bool{
	return s.top >= stackSize
}