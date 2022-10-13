package queue

import "errors"

type node[T comparable] struct {
	data T
	next *node[T]
}

type queue[T comparable] struct{
	count int
	head  *node[T]
	tail  *node[T]
}

// initialize new Queue
func NewQueue[T comparable]() *queue[T]{
	queue := new(queue[T])
	queue.count = 0
	queue.head = nil
	queue.tail = nil
	return queue
}

// create new node
func NewNode[T comparable](data T) *node[T]{
	node := new(node[T])
	node.data = data
	node.next = nil
	return node
}

// Enqueue - Add node to the end of the queue
func (q *queue[T]) Enqueue(data T) {
	newNode := NewNode(data)
	// if empty tail equal head
	if q.count == 0 {
		q.head = newNode
		q.tail = newNode
	} else{	
		// Point last node's next to the newnode and 
		// make newnode the tail
		q.tail.next = newNode
		q.tail = newNode
	}

	q.count++
}
// Dequeue - Remove node from the front of the queue
func ( q *queue[T] ) Dequeue() (*node[T], error) {
	// check if queue is empty
	if q.count == 0 {
		return nil, errors.New("queue is empty")
	}
	
	// save current head node
	current := q.head
	q.head = q.head.next
	q.count--

	return current, nil
}