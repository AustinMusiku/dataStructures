package queue

import "errors"

type node struct {
	data int
	next *node
}

type queue struct{
	count int
	head  *node
	tail  *node
}

// initialize new Queue
func NewQueue() *queue{
	return &queue{ 0, nil, nil}
}

// create new node
func NewNode(data int) *node{
	return &node{ data, nil }
}

// Enqueue - Add node to the end of the queue
func (q *queue) Enqueue(data int) {
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
func ( q *queue ) Dequeue() (*node, error) {
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