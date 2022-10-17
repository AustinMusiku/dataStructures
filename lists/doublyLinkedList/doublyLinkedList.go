package doublyLinkedList

import "errors"

type node[T comparable] struct {
	data T
	prev *node[T]
	next *node[T]
}

type doublyLinkedList[T comparable] struct {
	count int
	head  *node[T]
	tail  *node[T]
}

// initialize new Queue
func NewDList[T comparable]() *doublyLinkedList[T]{
	dList := new(doublyLinkedList[T])
	dList.count = 0
	dList.head = nil
	dList.tail = nil
	return dList
}

// create new node
func NewNode[T comparable](data T) *node[T]{
	node := new(node[T])
	node.data = data
	node.prev = nil
	node.next = nil
	return node
}

// AddBack - Add node at the end of the list
func (d *doublyLinkedList[T]) AddBack(data T) {
	newNode := NewNode(data)
	// if empty tail equal head
	if d.count == 0 {
		d.head = newNode
		d.tail = newNode
	} else{	
		// Point last node's next to the newnode
		d.tail.next = newNode
		// point newnode's prev to last node 
		newNode.prev = d.tail
		// make newnode the tail
		d.tail = newNode
	}

	d.count++
}

// AddFront - Add node at the front of the list
func (d *doublyLinkedList[T]) AddFront(data T) {
	newNode := NewNode(data)
	// if empty tail equal head
	if d.count == 0 {
		d.head = newNode
		d.tail = newNode
	} else{
		// Point newnode's next to the first and 
		newNode.next = d.head
		// point first node's prev to the newnode 
		d.head.prev = newNode
		// make newnode the head
		d.tail = newNode
	}

	d.count++
}

// InsertAt - Add node at a specific index
func ( d *doublyLinkedList[T] ) InsertAt(index int, data T) error {
	// out of bounds
	if index < 0 || index > d.count{
		return errors.New("out of bounds")
	}

	if d.IsEmpty() || index == d.count {
		// if empty or attempting to insert at the end
		d.AddBack(data)
	} else if index == 0 { 
		// else if attempting to insert at the front
		d.AddFront(data)
	} else {
		// traverse list till [index-1] and insert newnode  
		newNode := NewNode(data)
		current := d.head
		for i:=1; i<index; i++ {
			current = current.next
		}
		// Point newnode's next to the current's next and
		newNode.next = current.next
		// point current's next's prev to the newnode 
		current.next.prev = newNode
		// point current's next to the newnode 
		current.next = newNode
		// point newnode's prev to the current
		newNode.prev = current
	}

	d.count++
	return nil
}

// RemoveAt - Remove node at a specific index
func ( l *doublyLinkedList[T] ) RemoveAt(index int) ( *node[T], error ) {
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
		l.head.next.prev = nil
		l.count--
	} else {
		current := l.head
		for i:=1; i<index; i++ {
			current = current.next
		}
		removedNode = current.next
		current.next = current.next.next
		current.next.prev = current
		l.count--
	}

	return removedNode, nil
}

// Clear - Remove all nodes from the list
func ( d *doublyLinkedList[T] ) Clear() {
	d.count = 0
	d.head = nil
	d.tail = nil
}

// IsEmpty
func ( d *doublyLinkedList[T] ) IsEmpty() bool {
	return d.count == 0
}