package CircularBuffer

import (
	"errors"
	"reflect"
	"testing"
)

func TestCircularBuffer(t *testing.T) {
	t.Run("Create new buffer", func(t *testing.T) {
		buffer := NewCircularBuffer[int](5)
		if buffer.read != -1 {
			t.Errorf("Expected read to be -1, got %v", buffer.read)
		}
		if buffer.write != -1 {
			t.Errorf("Expected write to be -1, got %v", buffer.write)
		}
		if len(buffer.items) != 5 {
			t.Errorf("Expected buffer to have length 5, got %v", len(buffer.items))
		}
	})

	t.Run("Write to buffer", func(t *testing.T) {

		buffer := NewCircularBuffer[int](5)
		buffer.Write(1)

		if buffer.read != 0 {
			t.Errorf("Expected read to be 0, got %v", buffer.read)
		}
		if buffer.write != 0 {
			t.Errorf("Expected write to be 0, got %v", buffer.write)
		}
		if buffer.items[0] != 1 {
			t.Errorf("Expected buffer[0] to be 1, got %v", buffer.items[0])
		}
	})

	t.Run("Remove item from buffer", func(t *testing.T) {
		type testRead struct {
			args          []int
			expectedValue *int
			expectedError error
		}

		tests := []testRead{
			{[]int{}, nil, errors.New("buffer is empty. Can't read from buffer")},
			{[]int{1}, new(int), nil},
			{[]int{1, 2, 3, 4, 5}, new(int), nil},
		}

		for _, test := range tests {
			buffer := NewCircularBuffer[int](5)
			for _, arg := range test.args {
				buffer.Write(arg)
			}

			result, err := buffer.Read()

			if reflect.TypeOf(result) != reflect.TypeOf(test.expectedValue) {
				t.Errorf("Expected result to be of type %v, got %v",
					reflect.TypeOf(test.expectedValue),
					reflect.TypeOf(result))
			}

			if err != nil && test.expectedError == nil {
				t.Errorf("Expected error to be nil, got %v", err)
			}

			if err == nil && test.expectedError != nil {
				t.Errorf("Expected error to be %v, got nil", test.expectedError)
			}
		}
	})

	t.Run("Peek at buffer", func(t *testing.T) {
		buffer := NewCircularBuffer[int](5)
		buffer.Write(1)
		buffer.Write(2)
		buffer.Write(3)
		buffer.Write(4)
		buffer.Write(5)

		if buffer.Peek() != &buffer.items[0] {
			t.Errorf("Expected buffer[0] to be 1, got %v", buffer.items[0])
		}
	})

	t.Run("Resize buffer", func(t *testing.T) {
		buffer := NewCircularBuffer[int](5)
		buffer.Write(1)
		buffer.Write(2)
		buffer.Write(3)
		buffer.Write(4)
		buffer.Write(5)

		factor := 2
		buffer.Resize(2)

		if len(buffer.items) != 10 {
			t.Errorf("Expected buffer to have length %v, got %v",
				len(buffer.items)*factor,
				len(buffer.items))
		}

		if buffer.read != 0 {
			t.Errorf("Expected read to be 0, got %v", buffer.read)
		}

		if buffer.write != 4 {
			t.Errorf("Expected write to be 5, got %v", buffer.write)
		}
	})

	t.Run("Clear buffer", func(t *testing.T) {
		buffer := NewCircularBuffer[int](5)
		buffer.Write(1)
		buffer.Write(2)
		buffer.Write(3)
		buffer.Write(4)
		buffer.Write(5)

		clearedItems := buffer.Clear()

		if buffer.read != -1 {
			t.Errorf("Expected read to be -1, got %v", buffer.read)
		}
		if buffer.write != -1 {
			t.Errorf("Expected write to be -1, got %v", buffer.write)
		}
		if reflect.DeepEqual(clearedItems, []int{1, 2, 3, 4, 5}) != true {
			t.Errorf("Expected buffer clear to return %v, got %v",
				[]int{1, 2, 3, 4, 5},
				clearedItems)
		}
	})

	t.Run("Check if full", func(t *testing.T) {
		type testIsFull struct {
			args     []int
			expected bool
		}

		isFullTests := []testIsFull{
			{[]int{}, false},
			{[]int{1, 2, 3, 4}, false},
			{[]int{1, 2, 3, 4, 5}, true},
		}

		buffer := NewCircularBuffer[int](5)

		for _, test := range isFullTests {
			for _, arg := range test.args {
				buffer.Write(arg)
			}

			if buffer.IsFull() != test.expected {
				t.Errorf("Expected buffer to be full: %v, got %v",
					test.expected,
					buffer.IsFull())
			}
		}
	})

	t.Run("Check if empty", func(t *testing.T) {
		type testIsEmpty struct {
			args     []int
			expected bool
		}

		isEmptyTests := []testIsEmpty{
			{[]int{}, true},
			{[]int{1, 2, 3, 4}, false},
			{[]int{1, 2, 3, 4, 5}, false},
		}

		buffer := NewCircularBuffer[int](5)

		for _, test := range isEmptyTests {
			for _, arg := range test.args {
				buffer.Write(arg)
			}

			if buffer.IsEmpty() != test.expected {
				t.Errorf("Expected buffer to be empty: %v, got %v",
					test.expected,
					buffer.IsFull())
			}
		}
	})

}
