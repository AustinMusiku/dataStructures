package Heap

import "math"

type Sortable[T any] struct {
	Value    T
	Priority int
}

type Heap[T any] struct {
	items []Sortable[T]
	size  int
	mode  string
}

// Create a new heap
func NewHeap[T any](mode string) *Heap[T] {
	return &Heap[T]{
		items: make([]Sortable[T], 0),
		size:  0,
		mode:  mode,
	}
}

// Add item to the heap
func (h *Heap[T]) Insert(value T, priority int) {
	item := Sortable[T]{value, priority}

	// if no removals have been made, append to the end of the array
	// otherwise, insert at the end of the heap and shift the removed items 1 index to the right
	if h.size == len(h.items) {
		h.items = append(h.items, item)
	} else {
		newItems := make([]Sortable[T], len(h.items)+1)
		removed := h.items[h.size:]
		copy(newItems, h.items[:h.size])
		newItems[h.size] = item
		for i, x := range removed {
			newItems[(h.size+1)+i] = x
		}
		h.items = newItems
	}

	h.size++
	h.heapifyUp()
}

// Remove item from the heap.
// Returns the item with the highest priority
func (h *Heap[T]) Remove() Sortable[T] {
	current := h.items[0]

	h.swap(0, h.size-1)
	h.size--
	h.heapifyDown()
	return current
}

// Peek at the top item in the heap
func (h *Heap[T]) Peek() Sortable[T] {
	return h.items[0]
}

// Get the size of the heap
func (h *Heap[T]) Size() int {
	return h.size
}

// swap the values of two items in the heap
func (h *Heap[T]) swap(index1, index2 int) {
	h.items[index1], h.items[index2] = h.items[index2], h.items[index1]
}

// This method is called after a push to the heap.
// Moves the last inserted item up the heap to its correct position
func (h *Heap[T]) heapifyUp() {
	inserted := h.size - 1 // the index of the last inserted item (the last item in the heap)
	parent := getParentIndex(inserted)

	for h.hasParent(inserted) {
		// performSwap returns true if the inserted node has a higher priority than its parent
		performSwap := h.evaluateMode(h.items[inserted].Priority, h.items[parent].Priority)
		if performSwap {
			h.swap(inserted, parent)
		}
		inserted = parent
		parent = getParentIndex(inserted)
	}
}

// This method is called after a poll from the heap.
// Moves the item at the top down the heap to its correct position
func (h *Heap[T]) heapifyDown() {
	current := 0 // index of the first item in the heap
	child := -1

	for h.hasLeft(current) {
		// prioritiseRight is true if the right child is present and has a higher priority than the left child
		prioritiseRight := h.hasRight(current) && h.evaluateMode(h.getRight(current).Priority, h.getLeft(current).Priority)

		if prioritiseRight {
			child = getRightIndex(current)
		} else {
			child = getLeftIndex(current)
		}

		// perfomSwap is true if the current node has a child with a higher priority
		performSwap := h.evaluateMode(h.items[child].Priority, h.items[current].Priority)
		if performSwap {
			h.swap(current, child)
		}
		current = child
	}
}

// -------------------------------------------------------------------------------
// Helpers
// -------------------------------------------------------------------------------

// Returns the truthy value based on the mode of the heap
func (h *Heap[T]) evaluateMode(x, y int) bool {
	if h.mode == "max" {
		return x > y
	} else {
		return x < y
	}
}

// Returns true if the node has a parent
func (h *Heap[T]) hasParent(index int) bool {
	parentIndex := math.Floor((float64(index) - 1) / 2)
	return parentIndex >= 0
}

// Returns true if the node has a left child
func (h *Heap[T]) hasLeft(index int) bool {
	leftIndex := index*2 + 1
	return leftIndex < h.size
}

// Returns true if the node has a right child
func (h *Heap[T]) hasRight(index int) bool {
	leftIndex := index*2 + 2
	return leftIndex < h.size
}

// Returns the index of the parent
func getParentIndex(index int) int {
	return int(math.Floor((float64(index) - 1) / 2))
}

// Returns the index of the left child
func getLeftIndex(index int) int {
	return index*2 + 1
}

// Returns the index of the right child
func getRightIndex(index int) int {
	return index*2 + 2
}

// Returns the value of the left child
func (h *Heap[T]) getLeft(index int) *Sortable[T] {
	if h.hasLeft(index) {
		return &h.items[getLeftIndex(index)]
	}
	return nil
}

// Returns the value of the right child
func (h *Heap[T]) getRight(index int) *Sortable[T] {
	if h.hasRight(index) {
		return &h.items[getRightIndex(index)]
	}
	return nil
}
