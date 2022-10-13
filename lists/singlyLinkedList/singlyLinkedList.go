package singlyLinkedList

import "errors"

type node[T comparable] struct {
	data T
	next *node[T]
}

type singlyLinkedlist[T comparable] struct{
	count int
	head  *node[T]
	tail  *node[T]
}

// initialize new List
func NewList[T comparable]() *singlyLinkedlist[T]{
	list := new(singlyLinkedlist[T])
	list.count = 0
	list.head = nil
	list.tail = nil
	return list
}

// create new node
func NewNode[T comparable](data T) *node[T]{
	node := new(node[T])
	node.data = data
	node.next = nil
	return node
}


// Add - Add node at the end of the list
func ( l *singlyLinkedlist[T] ) Add(data T) {
	newNode := NewNode(data)
	// check if list is empty 
	if l.IsEmpty() {
		l.head = newNode
		l.tail = l.head
	} else{
		l.tail.next = newNode
		l.tail = newNode
	}
	l.count++
}

// InsertAt - Add node to a specific index
func ( l *singlyLinkedlist[T] ) InsertAt(index int, data T) error {
	// out of bounds
	if index < 0 || index > l.count{
		return errors.New("out of bounds")
	}
	// if empty or attempting to insert at the end
	if l.IsEmpty() || index == l.count {
		l.Add(data)
	} else {
		// traverse list till [index-1] and insert newnode  
		newNode := NewNode(data)
		current := l.head
		for i:=1; i<index; i++ {
			current = current.next
		}
		newNode.next = current.next
		current.next = newNode
	}

	l.count++
	return nil
}

// RemoveAt - Remove node at a specific index
func ( l *singlyLinkedlist[T] ) RemoveAt(index int) ( *node[T], error ) {
	// Check if out of bounds
	if index < 0 || index > l.count{
		return nil, errors.New("out of bounds")
	}
	
	var removedNode *node[T]

	if l.count == 1 {
		removedNode = l.head
		l.Clear()
	}else if index == 0 {
		removedNode = l.head
		l.head = l.head.next
		l.count--
	} else {
		current := l.head
		for i:=1; i<index; i++ {
			current = current.next
		}
		removedNode = current.next
		current.next = current.next.next
		l.count--
	}

	return removedNode, nil
}

// Clear - Remove all nodes from the list
func ( l *singlyLinkedlist[T] ) Clear() {
	l.count = 0
	l.head = nil
	l.tail = nil
}

// GetAt - Get an element at a specific index
func ( l *singlyLinkedlist[T] ) GetAt(index int) ( *node[T], error ){
	// Check if out of bounds
	if index < 0 || index > l.count{
		return nil, errors.New("out of bounds")
	}
	
	// Check if index is last node
	if index == l.count-1 {
		return l.tail, nil
	}

	current := l.head
	for i:=0; i<index; i++ {
		current = current.next
	}

	return current, nil
}


// IndexOf - Get Index of an element
func ( l *singlyLinkedlist[T] ) IndexOf(value T) int {
	index := -1
	for i, current := 0, l.head; current != nil; i, current = i+1, current.next {
		if current.data == value {
			index = i
			break
		}
	}
	return index
}

// IsEmpty
func ( l *singlyLinkedlist[T] ) IsEmpty() bool {
	return l.count == 0
}