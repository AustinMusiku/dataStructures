package stacksList

import "errors"

type Node struct {
	data int
	next *Node
}
type StackList struct { 
	count int
	top *Node
}

// insert data at the top of the stack(front of the list)
func (s *StackList) Push(data int){
	// create new node
	newNode := &Node{ data: data }

	// check for stack underflow
	if s.count == 0 {
		s.top = newNode
	} else{
		newNode.next = s.top
		s.top = newNode
	}
	s.count++
}

// remove data at the top of the stack(front of the list)
func (s *StackList) Pop() (*Node, error) {
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
func (s *StackList) Peek()  (*Node, error){
	// check for stack underflow
	if s.count == 0{
		return nil, errors.New("stack is empty")
	}
	return s.top, nil
}