package stackList

import "errors"

type node[T any] struct {
	data T
	next *node[T]
}
type stackList[T any] struct {
	count int
	top   *node[T]
}

// initialize new List
func NewList[T any]() *stackList[T] {
	list := new(stackList[T])
	list.count = 0
	list.top = nil
	return list
}

// create new node
func NewNode[T any](data T) *node[T] {
	node := new(node[T]) 
	node.data = data 
	node.next = nil 
	return node
}

// insert data at the top of the stack(front of the list)
func (s *stackList[T]) Push(data T) {
	// create new node
	newNode := NewNode(data)

	// check for stack underflow
	if s.count == 0 {
		s.top = newNode
	} else {
		newNode.next = s.top
		s.top = newNode
	}
	s.count++
}

// remove data at the top of the stack(front of the list)
func (s *stackList[T]) Pop() (*node[T], error) {
	// check for stack underflow
	if s.count == 0 {
		return nil, errors.New("stack is empty")
	}

	current := s.top
	s.top = s.top.next
	s.count--
	return current, nil
}

// peek data at the top of the stack(front of the list)
func (s *stackList[T]) Peek() (*node[T], error) {
	// check for stack underflow
	if s.count == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.top, nil
}
