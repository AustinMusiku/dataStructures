package Heap

import "testing"

func TestHeap(t *testing.T) {
	t.Parallel()

	modes := []string{"max", "min"}

	t.Run("Create a heap", func(t *testing.T) {
		t.Parallel()

		for _, mode := range modes {
			heap := NewHeap[int](mode)

			if heap.mode != mode {
				t.Errorf("Expected heap mode to be %s, got %s", mode, heap.mode)
			}

			if heap.Size() != 0 {
				t.Errorf("Expected heap size to be 0, got %d", heap.Size())
			}
		}
	})

	t.Run("Insert to heap", func(t *testing.T) {
		t.Parallel()

		type testInsert struct {
			args         []Sortable[int]
			expectedTop  []Sortable[int]
			expectedLast []Sortable[int]
		}

		testCase := testInsert{
			args: []Sortable[int]{
				{Value: 8, Priority: 8},
				{Value: 9, Priority: 9},
				{Value: 12, Priority: 12},
				{Value: 7, Priority: 7},
				{Value: 11, Priority: 11},
			},
			expectedTop: []Sortable[int]{
				{Value: 12, Priority: 12},
				{Value: 7, Priority: 7},
			},
			expectedLast: []Sortable[int]{
				{Value: 8, Priority: 8},
				{Value: 11, Priority: 11},
			},
		}

		for i, mode := range modes {

			heap := NewHeap[int](mode)

			for _, arg := range testCase.args {
				heap.Insert(arg.Value, arg.Priority)
			}

			foundTop := heap.Peek()
			expectedTop := testCase.expectedTop[i]

			foundLastItem := heap.items[heap.Size()-1]
			expectedLastItem := testCase.expectedLast[i]

			if heap.Size() != len(testCase.args) {
				t.Errorf("Expected heap size to be %d, got %d", len(testCase.args), heap.Size())
			}

			if mode == heap.mode && foundTop.Priority != expectedTop.Priority {
				t.Errorf("Expected %s top item to be %d, got %d", mode, expectedTop.Priority, foundTop.Priority)
			}

			if mode == heap.mode && foundLastItem.Priority != expectedLastItem.Priority {
				t.Errorf("Expected %s last item to be %d, got %d", mode, expectedLastItem.Priority, foundLastItem.Priority)
			}
		}
	})

	t.Run("Remove from heap", func(t *testing.T) {
		t.Parallel()

		type testRemove struct {
			args            []Sortable[int]
			expectedRemoved []Sortable[int]
			expectedTop     []Sortable[int]
			expectedLast    []Sortable[int]
		}

		testCase := testRemove{
			args: []Sortable[int]{
				{Value: 8, Priority: 8},
				{Value: 9, Priority: 9},
				{Value: 12, Priority: 12},
				{Value: 7, Priority: 7},
				{Value: 11, Priority: 11},
			},
			expectedRemoved: []Sortable[int]{
				{Value: 12, Priority: 12},
				{Value: 7, Priority: 7},
			},
			expectedTop: []Sortable[int]{
				{Value: 11, Priority: 11},
				{Value: 8, Priority: 8},
			},
			expectedLast: []Sortable[int]{
				{Value: 7, Priority: 7},
				{Value: 11, Priority: 11},
			},
		}

		for i, mode := range modes {

			heap := NewHeap[int](mode)

			for _, arg := range testCase.args {
				heap.Insert(arg.Value, arg.Priority)
			}

			removed := heap.Remove()
			expectedRemoved := testCase.expectedRemoved[i]

			foundTop := heap.Peek()
			expectedTop := testCase.expectedTop[i]

			foundLastItem := heap.items[heap.Size()-1]
			expectedLastItem := testCase.expectedLast[i]

			if heap.Size() != len(testCase.args)-1 {
				t.Errorf("Expected heap size to be %d, got %d", len(testCase.args)-1, heap.Size())
			}

			if removed.Priority != expectedRemoved.Priority {
				t.Errorf("Expected removed item to be %d, got %d", expectedRemoved.Priority, removed.Priority)
			}

			if mode == heap.mode && foundTop.Priority != expectedTop.Priority {
				t.Errorf("Expected %s top item to be %d, got %d", mode, expectedTop.Priority, foundTop.Priority)
			}

			if mode == heap.mode && foundLastItem.Priority != expectedLastItem.Priority {
				t.Errorf("Expected %s last item to be %d, got %d", mode, expectedLastItem.Priority, foundLastItem.Priority)
			}
		}
	})

	t.Run("Peek at heap", func(t *testing.T) {
		t.Parallel()

		type testPeek struct {
			args     []Sortable[int]
			expected []Sortable[int]
		}

		testCase := testPeek{
			args: []Sortable[int]{
				{Value: 8, Priority: 8},
				{Value: 9, Priority: 9},
				{Value: 12, Priority: 12},
				{Value: 7, Priority: 7},
				{Value: 11, Priority: 11},
			},
			expected: []Sortable[int]{
				{Value: 12, Priority: 12},
				{Value: 7, Priority: 7},
			},
		}

		for i, mode := range modes {

			heap := NewHeap[int](mode)

			for _, arg := range testCase.args {
				heap.Insert(arg.Value, arg.Priority)
			}

			found := heap.Peek()
			expected := testCase.expected[i]

			if found.Priority != expected.Priority {
				t.Errorf("Expected %s heap peeked item to be %d, got %d", mode, expected.Priority, found.Priority)
			}
		}
	})

	t.Run("Get size of heap", func(t *testing.T) {
		t.Parallel()

		heap := NewHeap[int]("max")

		heap.Insert(1, 1)

		if heap.Size() != 1 {
			t.Errorf("Expected heap size to be 1, got %d", heap.Size())
		}
	})
}
