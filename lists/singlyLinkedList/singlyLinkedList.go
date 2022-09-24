package singlyLinkedList

import "errors"

type node struct {
	data int
	next *node
}

type singlyLinkedlist struct{
	count int
	head  *node
	tail  *node
}

// initialize new List
func NewList() *singlyLinkedlist{
	return &singlyLinkedlist{ 0, nil, nil }
}

// create new node
func NewNode(data int) *node{
	return &node{ data, nil }
}


// Add. Add node at the end of the list
func ( l *singlyLinkedlist ) Add(data int) {
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

// InsertAt. Add node to a specific index
func ( l *singlyLinkedlist ) InsertAt(index, data int) error {
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

// RemoveAt. Remove node at a specific index
func ( l *singlyLinkedlist ) RemoveAt(index int) ( *node, error ) {
	// Check if out of bounds
	if index < 0 || index > l.count{
		return nil, errors.New("out of bounds")
	}
	
	var removedNode *node

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

// Clear. Remove all nodes from the list
func ( l *singlyLinkedlist ) Clear() {
	l.count = 0
	l.head = nil
	l.tail = nil
}

// GetAt. Get an element at a specific index
func ( l *singlyLinkedlist ) GetAt(index int) ( *node, error ){
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


// IndexOf. Get Index of an element
func ( l *singlyLinkedlist ) IndexOf(value int) int {
	index := -1
	for i, current := 0, l.head; current != nil; i, current = i+1, current.next {
		if current.data == value {
			index = i
		}
	}
	return index
}

// IsEmpty
func ( l *singlyLinkedlist ) IsEmpty() bool {
	return l.count == 0
}