package HashMap

import (
	"errors"
	"testing"
)

type nodeType struct {
	key   int
	value string
}

func TestHashMapBucket(t *testing.T) {
	t.Parallel()

	testNodes := []nodeType{
		{1, "Rick Grimes"},
		{2, "Daryl Dixon"},
		{3, "Glenn Rhee"},
	}

	t.Run("Create new bucket", func(t *testing.T) {
		bucket := NewBucket[int, string]()

		if bucket.count != 0 {
			t.Errorf("Expected count to be 0, got %d", bucket.count)
		}

		if bucket.head != nil {
			t.Errorf("Expected head to be nil, got %v", bucket.head)
		}
	})

	t.Run("Add new node to bucket", func(t *testing.T) {
		t.Parallel()

		bucket := NewBucket[int, string]()

		for _, node := range testNodes {
			bucket.Add(node.key, node.value)
		}

		if bucket.count != 3 {
			t.Errorf("Expected count to be 3, got %d", bucket.count)
		}

		if bucket.head == nil {
			t.Errorf("Expected head to not be nil, got %v", bucket.head)
		}

		if bucket.head.key != 3 {
			t.Errorf("Expected head key to be 3, got %v", bucket.head.key)
		}

		if bucket.head.value != "Glenn Rhee" {
			t.Errorf("Expected head value to be Glenn Rhee, got %v", bucket.head.value)
		}
	})

	t.Run("Get node from bucket", func(t *testing.T) {
		t.Parallel()

		type testGet struct {
			data         []nodeType
			getKey       int
			expectedNode *node[int, string]
			expectedErr  error
		}

		testCases := []testGet{
			{
				data:         testNodes,
				getKey:       2,
				expectedNode: NewNode[int, string](2, "Daryl Dixon"),
				expectedErr:  nil,
			},
			{
				data:         testNodes,
				getKey:       4,
				expectedNode: nil,
				expectedErr:  errors.New("key does not exist in map"),
			},
		}

		for _, tc := range testCases {
			bucket := NewBucket[int, string]()

			for _, node := range tc.data {
				bucket.Add(node.key, node.value)
			}

			node, err := bucket.Get(tc.getKey)

			if node != nil {
				// positive test case
				if node.key != tc.expectedNode.key {
					t.Errorf("Expected node key to be %d, got %d", tc.expectedNode.key, node.key)
				}

				if node.value != tc.expectedNode.value {
					t.Errorf("Expected node value to be %s, got %s", tc.expectedNode.value, node.value)
				}

				if err != nil {
					t.Errorf("Expected err to be nil, got %v", err)
				}

			} else {
				// negative test case
				if tc.expectedNode != nil {
					t.Errorf("Expected node to not be nil, got %v", node)
				}

				if err.Error() != tc.expectedErr.Error() {
					t.Errorf("Expected err to be %v, got %v", tc.expectedErr, err)
				}
			}
		}
	})

	t.Run("Update node in bucket", func(t *testing.T) {
		t.Parallel()

		bucket := NewBucket[int, string]()

		for _, node := range testNodes {
			bucket.Add(node.key, node.value)
		}

		bucket.Update(2, "Carl Grimes")

		node, err := bucket.Get(2)

		if node == nil {
			t.Errorf("Expected node to not be nil, got %v", node)
		}

		if err != nil {
			t.Errorf("Expected err to be nil, got %v", err)
		}

		if node.key != 2 {
			t.Errorf("Expected node key to be 2, got %v", node.key)
		}

		if node.value != "Carl Grimes" {
			t.Errorf("Expected node value to be Carl Grimes, got %v", node.value)
		}
	})

	t.Run("Remove node from bucket", func(t *testing.T) {
		t.Parallel()

		type testRemove struct {
			data         []nodeType
			removeKey    int
			expectedNode *node[int, string]
			expectedErr  error
		}

		testCases := []testRemove{
			{
				data:         testNodes,
				removeKey:    2,
				expectedNode: NewNode[int, string](2, "Daryl Dixon"),
				expectedErr:  nil,
			},
			{
				data:         testNodes,
				removeKey:    4,
				expectedNode: nil,
				expectedErr:  errors.New("key does not exist in map"),
			},
		}

		for _, test := range testCases {
			bucket := NewBucket[int, string]()

			for _, node := range test.data {
				bucket.Add(node.key, node.value)
			}

			removed, err := bucket.Remove(test.removeKey)

			if removed != nil {
				// positive test case
				if removed.key != test.expectedNode.key {
					t.Errorf("Expected removed key to be %v, got %v", test.expectedNode.key, removed.key)
				}

				if removed.value != test.expectedNode.value {
					t.Errorf("Expected removed value to be %v, got %v", test.expectedNode.value, removed.value)
				}
			} else {
				// negative test case
				if test.expectedNode != nil {
					t.Errorf("Expected removed node to not be nil, got %v", removed)
				}

				if err == nil {
					t.Errorf("Expected err to not be nil, got %v", err)
				}

				if err.Error() != test.expectedErr.Error() {
					t.Errorf("Expected err to be %v, got %v", test.expectedErr.Error(), err.Error())
				}
			}
		}
	})

	t.Run("Check if node exists in bucket", func(t *testing.T) {
		t.Parallel()

		type testContains struct {
			data     []nodeType
			key      int
			expected int
		}

		testCases := []testContains{
			{
				data:     testNodes,
				key:      2,
				expected: 1,
			},
			{
				data:     testNodes,
				key:      4,
				expected: -1,
			},
		}

		for _, test := range testCases {
			bucket := NewBucket[int, string]()

			for _, node := range test.data {
				bucket.Add(node.key, node.value)
			}

			contains := bucket.Contains(test.key)

			if contains != test.expected {
				t.Errorf("Expected contains to be %v, got %v", test.expected, contains)
			}
		}
	})

	t.Run("Clear bucket", func(t *testing.T) {
		t.Parallel()

		bucket := NewBucket[int, string]()

		for _, node := range testNodes {
			bucket.Add(node.key, node.value)
		}

		bucket.Clear()

		if bucket.count != 0 {
			t.Errorf("Expected count to be 0, got %d", bucket.count)
		}

		if bucket.head != nil {
			t.Errorf("Expected head to be nil, got %v", bucket.head)
		}
	})

}
