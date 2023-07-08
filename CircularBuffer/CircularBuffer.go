package CircularBuffer

import (
	"errors"
)

type circularBuffer[T any] struct {
	items []T
	read  int
	write int
}

func NewCircularBuffer[T any](size uint) *circularBuffer[T] {
	return &circularBuffer[T]{
		items: make([]T, size),
		read:  -1,
		write: -1,
	}
}

// Add items to the buffer
func (c *circularBuffer[T]) Write(data T) error {
	if c.IsFull() {
		return errors.New("buffer is full. Can't write to buffer")
	}
	if c.IsEmpty() {
		c.read++
	}

	c.write = (c.write + 1) % len(c.items)
	c.items[c.write] = data
	return nil
}

// Consume items from the buffer
func (c *circularBuffer[T]) Read() (*T, error) {
	if c.IsEmpty() {
		return nil, errors.New("buffer is empty. Can't read from buffer")
	}

	current := c.items[c.read]

	if c.write == c.read {
		c.write--
		c.read--
	}

	c.read = (c.read + 1) % len(c.items)
	return &current, nil
}

// Return the next item to be consumed
func (c *circularBuffer[T]) Peek() *T {
	return &c.items[c.read]
}

func (c *circularBuffer[T]) Resize(factor uint) error {
	tempArr := make([]T, len(c.items)*int(factor))

	if elementsCopied := copy(tempArr, c.items); elementsCopied != len(c.items) {
		return errors.New("failed to resize buffer. Couldn't copy elements")
	}

	c.items = tempArr
	return nil
}

// Reset the buffer and return the items currently in the buffer
func (c *circularBuffer[T]) Clear() []T {
	c.read = -1
	c.write = -1
	return c.items
}

func (c *circularBuffer[T]) IsFull() bool {
	nextPos := (c.write + 1) % len(c.items)
	return nextPos == c.read
}

func (c *circularBuffer[T]) IsEmpty() bool {
	return (c.read == -1 && c.write == -1)
}
