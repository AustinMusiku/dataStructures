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
				{value: 8, priority: 8},
				{value: 9, priority: 9},
				{value: 12, priority: 12},
				{value: 7, priority: 7},
				{value: 11, priority: 11},
			},
			expectedTop: []Sortable[int]{
				{value: 12, priority: 12},
				{value: 7, priority: 7},
			},
			expectedLast: []Sortable[int]{
				{value: 8, priority: 8},
				{value: 11, priority: 11},
			},
		}

		for i, mode := range modes {

			heap := NewHeap[int](mode)

			for _, arg := range testCase.args {
				heap.Insert(arg.value, arg.priority)
			}

			foundTop := heap.Peek()
			expectedTop := testCase.expectedTop[i]

			foundLastItem := heap.items[heap.Size()-1]
			expectedLastItem := testCase.expectedLast[i]

			if heap.Size() != len(testCase.args) {
				t.Errorf("Expected heap size to be %d, got %d", len(testCase.args), heap.Size())
			}

			if mode == heap.mode && foundTop.priority != expectedTop.priority {
				t.Errorf("Expected %s top item to be %d, got %d", mode, expectedTop.priority, foundTop.priority)
			}

			if mode == heap.mode && foundLastItem.priority != expectedLastItem.priority {
				t.Errorf("Expected %s last item to be %d, got %d", mode, expectedLastItem.priority, foundLastItem.priority)
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
				{value: 8, priority: 8},
				{value: 9, priority: 9},
				{value: 12, priority: 12},
				{value: 7, priority: 7},
				{value: 11, priority: 11},
			},
			expectedRemoved: []Sortable[int]{
				{value: 12, priority: 12},
				{value: 7, priority: 7},
			},
			expectedTop: []Sortable[int]{
				{value: 11, priority: 11},
				{value: 8, priority: 8},
			},
			expectedLast: []Sortable[int]{
				{value: 7, priority: 7},
				{value: 11, priority: 11},
			},
		}

		for i, mode := range modes {

			heap := NewHeap[int](mode)

			for _, arg := range testCase.args {
				heap.Insert(arg.value, arg.priority)
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

			if removed.priority != expectedRemoved.priority {
				t.Errorf("Expected removed item to be %d, got %d", expectedRemoved.priority, removed.priority)
			}

			if mode == heap.mode && foundTop.priority != expectedTop.priority {
				t.Errorf("Expected %s top item to be %d, got %d", mode, expectedTop.priority, foundTop.priority)
			}

			if mode == heap.mode && foundLastItem.priority != expectedLastItem.priority {
				t.Errorf("Expected %s last item to be %d, got %d", mode, expectedLastItem.priority, foundLastItem.priority)
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
				{value: 8, priority: 8},
				{value: 9, priority: 9},
				{value: 12, priority: 12},
				{value: 7, priority: 7},
				{value: 11, priority: 11},
			},
			expected: []Sortable[int]{
				{value: 12, priority: 12},
				{value: 7, priority: 7},
			},
		}

		for i, mode := range modes {

			heap := NewHeap[int](mode)

			for _, arg := range testCase.args {
				heap.Insert(arg.value, arg.priority)
			}

			found := heap.Peek()
			expected := testCase.expected[i]

			if found.priority != expected.priority {
				t.Errorf("Expected %s heap peeked item to be %d, got %d", mode, expected.priority, found.priority)
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
