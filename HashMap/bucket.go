package HashMap

import (
	"errors"
	"sync"
)

type node[K Hashable, V any] struct {
	mu    sync.Mutex
	key   K
	value V
	next  *node[K, V]
}

type bucket[K Hashable, V any] struct {
	mu    sync.Mutex
	count int
	head  *node[K, V]
}

// initialize new bucket
func NewBucket[K Hashable, V any]() *bucket[K, V] {
	bucket := new(bucket[K, V])
	bucket.count = 0
	bucket.head = nil
	return bucket
}

// initialize new node
func NewNode[K Hashable, V any](key K, value V) *node[K, V] {
	node := new(node[K, V])
	node.key = key
	node.value = value
	node.next = nil
	return node
}

// Add - Add node at the front of the bucket
func (b *bucket[K, V]) Add(key K, value V) {
	b.mu.Lock()
	defer b.mu.Unlock()

	newNode := NewNode(key, value)
	// check if bucket is empty
	if b.IsEmpty() {
		b.head = newNode
	} else {
		newNode.next = b.head
		b.head = newNode
	}
	b.count++
}

// Get - get the value of a node by key
func (b *bucket[K, V]) Get(key K) (*node[K, V], error) {
	for current := b.head; current != nil; current = current.next {
		// if key matches
		if current.key == key {
			return current, nil
		}
	}

	return nil, errors.New("key does not exist in map")
}

// Update - modify the value of a node by key
func (b *bucket[K, V]) Update(key K, value V) (*node[K, V], error) {
	for current := b.head; current != nil; current = current.next {
		// if key matches
		if current.key == key {
			current.mu.Lock()
			current.value = value
			current.mu.Unlock()

			return current, nil
		}
	}

	return nil, errors.New("key does not exist in map")
}

// Remove - Delete a node by key
func (b *bucket[K, V]) Remove(key K) (*node[K, V], error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	var deleted *node[K, V] = nil

	for prev, current := b.head, b.head; current != nil; current = current.next {
		// if key matches
		if current.key == key {
			deleted = current

			if current == b.head {
				b.head = current.next
			} else {
				prev.next = current.next
			}

			b.count--
			return deleted, nil
		}

		// skip prev-node update on first iteration
		if current != b.head {
			prev = prev.next
		}
	}

	return deleted, errors.New("key does not exist in map")
}

// Contains - Check if a key exists in the bucket.
// Return index if found, -1 otherwise
func (b *bucket[K, V]) Contains(key K) int {
	for i, current := 0, b.head; current != nil; i, current = i+1, current.next {
		if current.key == key {
			return i
		}
	}
	return -1
}

// Clear - Remove all nodes from the bucket
func (b *bucket[K, V]) Clear() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.count = 0
	b.head = nil
}

// IsEmpty
func (b *bucket[K, V]) IsEmpty() bool {
	return b.count == 0
}
