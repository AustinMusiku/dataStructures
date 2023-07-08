package singlyLinkedList

import (
	"errors"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	t.Parallel()

	t.Run("Create new list", func(t *testing.T) {
		t.Parallel()
		list := NewList[int]()
		if list.count != 0 {
			t.Errorf("Expected count to be 0, got %v", list.count)
		}
		if list.head != nil {
			t.Errorf("Expected head to be nil, got %v", list.head)
		}
		if list.tail != nil {
			t.Errorf("Expected tail to be nil, got %v", list.tail)
		}
	})

	t.Run("Add item to list", func(t *testing.T) {
		t.Parallel()
		list := NewList[int]()
		list.Add(1)
		list.Add(2)
		list.Add(3)

		if list.count != 3 {
			t.Errorf("Expected count to be 3, got %v", list.count)
		}
		if list.head.data != 1 {
			t.Errorf("Expected head to hold 1, got %v", list.head.data)
		}
		if list.head.next.data != 2 {
			t.Errorf("Expected head->next to hold 2, got %v", list.head.next.data)
		}
		if list.tail.data != 3 {
			t.Errorf("Expected tail to hold 3, got %v", list.tail.data)
		}
	})

	t.Run("Insert item at index", func(t *testing.T) {
		t.Parallel()
		list := NewList[int]()
		list.Add(1)
		list.Add(2)
		list.Add(3)
		list.InsertAt(1, 4)

		if list.count != 4 {
			t.Errorf("Expected count to be 4, got %v", list.count)
		}
		if list.head.data != 1 {
			t.Errorf("Expected head to hold 1, got %v", list.head.data)
		}
		if list.head.next.data != 4 {
			t.Errorf("Expected head->next to hold 4, got %v", list.head.next.data)
		}
		if list.head.next.next.data != 2 {
			t.Errorf("Expected head->next->next to hold 2, got %v", list.head.next.next.data)
		}
		if list.tail.data != 3 {
			t.Errorf("Expected tail to hold 3, got %v", list.tail.data)
		}
	})

	t.Run("Remove item at index", func(t *testing.T) {
		t.Parallel()
		type testRemoveAt struct {
			args          []int
			targetIndex   int
			expectedValue int
			expectedError error
		}

		tests := []testRemoveAt{
			{[]int{}, 2, 0, errors.New("out of bounds")},
			{[]int{1}, 0, 1, nil},
			{[]int{1, 2, 3}, 0, 1, nil},
			{[]int{1, 2, 3}, 1, 2, nil},
			{[]int{1, 2, 3}, 2, 3, nil},
		}

		for _, test := range tests {
			list := NewList[int]()
			for _, val := range test.args {
				list.Add(val)
			}

			value, err := list.RemoveAt(test.targetIndex)

			if err != nil && test.expectedError == nil {
				t.Errorf("Expected error to be nil, got %v", err)
			}

			if err == nil && test.expectedError != nil {
				t.Errorf("Expected error to be %v, got nil", test.expectedError)
			}

			if err == nil {
				got := value.data
				want := test.expectedValue
				if got != want {
					t.Errorf("Expected value to be %v, got %v", want, got)
				}
			}

			if test.expectedError == nil && list.count != len(test.args)-1 {
				t.Errorf("Expected count to be %v, got %v", len(test.args)-1, list.count)
			}
		}
	})

	t.Run("Clear list", func(t *testing.T) {
		t.Parallel()
		list := NewList[int]()
		list.Add(1)
		list.Add(2)
		list.Add(3)
		list.Clear()

		if list.count != 0 {
			t.Errorf("Expected count to be 0, got %v", list.count)
		}
		if list.head != nil {
			t.Errorf("Expected head to be nil, got %v", list.head)
		}
		if list.tail != nil {
			t.Errorf("Expected tail to be nil, got %v", list.tail)
		}
	})

	t.Run("Get item at index", func(t *testing.T) {
		t.Parallel()
		type testGetAt struct {
			args          []int
			targetIndex   int
			expectedValue int
			expectedError error
		}

		tests := []testGetAt{
			{[]int{}, 2, 0, errors.New("out of bounds")},
			{[]int{1}, 0, 1, nil},
			{[]int{1, 2, 3}, 0, 1, nil},
			{[]int{1, 2, 3}, 1, 2, nil},
			{[]int{1, 2, 3}, 2, 3, nil},
		}

		for _, test := range tests {
			list := NewList[int]()
			for _, val := range test.args {
				list.Add(val)
			}

			value, err := list.GetAt(test.targetIndex)

			if err != nil && test.expectedError == nil {
				t.Errorf("Expected error to be nil, got %v", err)
			}

			if err == nil && test.expectedError != nil {
				t.Errorf("Expected error to be %v, got nil", test.expectedError)
			}

			if err == nil {
				got := value.data
				want := test.expectedValue

				if got != want {
					t.Errorf("Expected value to be %v, got %v", test.expectedValue, value)
				}
			}

		}
	})

	t.Run("Get index of item", func(t *testing.T) {
		t.Parallel()
		list := NewList[int]()
		list.Add(1)
		list.Add(2)
		list.Add(3)

		type testIndexOf struct {
			args          []int
			targetValue   int
			expectedValue int
		}

		tests := []testIndexOf{
			{[]int{}, 2, -1},
			{[]int{1}, 1, 0},
			{[]int{1, 2, 3}, 2, 1},
			{[]int{1, 2, 3}, 3, 2},
		}

		for _, test := range tests {
			list := NewList[int]()
			for _, val := range test.args {
				list.Add(val)
			}

			value := list.IndexOf(test.targetValue)

			if value != test.expectedValue {
				t.Errorf("Expected value to be %v, got %v", test.expectedValue, value)
			}
		}
	})
}
