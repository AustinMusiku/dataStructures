package singlyLinkedList

import (
	"errors"
	"reflect"
	"sync"
)

type Node[T any] struct {
	mu   sync.Mutex
	data T
	next *Node[T]
}

type SinglyLinkedlist[T any] struct {
	mu    sync.Mutex
	count int
	head  *Node[T]
	tail  *Node[T]
}

// initialize new List
func NewList[T any]() *SinglyLinkedlist[T] {
	list := new(SinglyLinkedlist[T])
	list.count = 0
	list.head = nil
	list.tail = nil
	return list
}

// create new node
func NewNode[T any](data T) *Node[T] {
	node := new(Node[T])
	node.data = data
	node.next = nil
	return node
}

func (n *Node[T]) GetData() T {
	return n.data
}

// --------------------------------------------
// GETTERS ------------------------------------
// --------------------------------------------
func (n *Node[T]) GetNext() *Node[T] {
	return n.next
}

func (l *SinglyLinkedlist[T]) GetCount() int {
	return l.count
}

func (l *SinglyLinkedlist[T]) GetHead() *Node[T] {
	return l.head
}

func (l *SinglyLinkedlist[T]) GetTail() *Node[T] {
	return l.tail
}

// --------------------------------------------
// SETTERS ------------------------------------
// --------------------------------------------
func (n *Node[T]) SetData(data T) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.data = data
}

func (n *Node[T]) SetNext(next *Node[T]) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.next = next
}

func (l *SinglyLinkedlist[T]) SetCount(count int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.count = count
}

func (l *SinglyLinkedlist[T]) SetHead(head *Node[T]) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.head = head
}

func (l *SinglyLinkedlist[T]) SetTail(tail *Node[T]) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.tail = tail
}

// --------------------------------------------
// LIST METHODS -------------------------------
// --------------------------------------------

// Add - Add node at the end of the list
func (l *SinglyLinkedlist[T]) Add(data T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	newNode := NewNode(data)
	// check if list is empty
	if l.IsEmpty() {
		l.head = newNode
		l.tail = l.head
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.count++
}

// InsertAt - Add node to a specific index
func (l *SinglyLinkedlist[T]) InsertAt(index int, data T) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// out of bounds
	if index < 0 || index > l.count {
		return errors.New("out of bounds")
	}
	// if empty or attempting to insert at the end
	if l.IsEmpty() || index == l.count {
		l.Add(data)
		return nil
	}

	newNode := NewNode(data)

	// if inserting at the beginning
	if index == 0 {
		newNode.next = l.head
		l.head = newNode
	} else {
		current := l.head
		for i := 1; i < index; i++ {
			current = current.next
		}
		newNode.next = current.next
		current.next = newNode
	}

	l.count++
	return nil
}

// RemoveAt - Remove node at a specific index
func (l *SinglyLinkedlist[T]) RemoveAt(index int) (*Node[T], error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Check if out of bounds
	if index < 0 || index > l.count {
		return nil, errors.New("out of bounds")
	}

	var removedNode *Node[T]

	if l.count == 1 {
		removedNode = l.head
		l.head = nil
		l.tail = nil
		l.count--
	} else if index == 0 {
		removedNode = l.head
		l.head = l.head.next
		l.count--
	} else {
		current := l.head
		for i := 1; i < index; i++ {
			current = current.next
		}
		removedNode = current.next
		current.next = current.next.next
		l.count--
	}

	return removedNode, nil
}

// Clear - Remove all nodes from the list
func (l *SinglyLinkedlist[T]) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.count = 0
	l.head = nil
	l.tail = nil
}

// GetAt - Get an element at a specific index
func (l *SinglyLinkedlist[T]) GetAt(index int) (*Node[T], error) {
	// Check if out of bounds
	if index < 0 || index > l.count {
		return nil, errors.New("out of bounds")
	}

	// Check if index is last node
	if index == l.count-1 {
		return l.tail, nil
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current, nil
}

// IndexOf - Get Index of an element
func (l *SinglyLinkedlist[T]) IndexOf(value T) int {
	index := -1
	for i, current := 0, l.head; current != nil; i, current = i+1, current.next {
		if reflect.DeepEqual(current.data, value) {
			index = i
			break
		}
	}
	return index
}

// IsEmpty
func (l *SinglyLinkedlist[T]) IsEmpty() bool {
	return l.count == 0
}
