package dQueue

import "errors"

type node[T comparable] struct {
	data T
	next *node[T]
}

type dQueue[T comparable] struct {
	count int
	head  *node[T]
	tail  *node[T]
}

// initialize new Queue
func NewQueue[T comparable]() *dQueue[T]{
	dQueue := new(dQueue[T])
	dQueue.count = 0
	dQueue.head = nil
	dQueue.tail = nil
	return dQueue
}

// create new node
func NewNode[T comparable](data T) *node[T]{
	node := new(node[T])
	node.data = data
	node.next = nil
	return node
}

// EnqueueFront - Reverse enqueue.
// Add node to the front of the queue.
func (q *dQueue[T]) EnqueueFront(data T) {
	newNode := NewNode(data)
	// if empty tail equal head
	if q.count == 0 {
		q.head = newNode
		q.tail = newNode
	} else{	
		// Point new node's next to the current first node and 
		// point the head at the new node
		newNode.next = q.head
		q.head = newNode
	}

	q.count++
}

// EnqueueBack.
// Add node to the Back of the queue.
func (q *dQueue[T]) EnqueueBack(data T) {
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

// DequeueFront.
// Remove node from the front of the queue.
func ( q *dQueue[T] ) DequeueFront() (*node[T], error) {
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

// DequeueBack - Reverse dequeue.  
// Remove node from the Back of the queue.
func ( q *dQueue[T] ) DequeueBack() (*node[T], error) {
	// check if queue is empty
	if q.count == 0 {
		return nil, errors.New("queue is empty")
	}
	
	// save current head node
	current := q.tail
	
	// traverse to the second last item and set it as the new tail
	secondLast, err := q.Traverse(q.count - 2)

	if err != nil {
		return nil, errors.New("error removing from end of queue")
	}

	q.tail = secondLast
	q.count--

	return current, nil
}

// traverse
func (q *dQueue[T]) Traverse(i int) (*node[T] ,error) {
	// check out of Bounds
	if i >= q.count {
		return nil, errors.New("out of Bounds")	
	}

	current := q.head

	for j:=0; j<i; j++{
		current = current.next
	}
	return current, nil
}