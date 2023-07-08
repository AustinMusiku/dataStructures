package AVLTree

import (
	"testing"
)

func TestAVLTree(t *testing.T) {

	t.Run("Create new tree", func(t *testing.T) {
		tree := NewAVLTree[int]()
		if tree.Root != nil {
			t.Errorf("Expected root to be nil, got %v", tree.Root)
		}
	})

	t.Run("Create new node", func(t *testing.T) {
		node := NewNode[int](5)
		if node.Data != 5 {
			t.Errorf("Expected node data to be 5, got %v", node.Data)
		}
		if node.Left != nil {
			t.Errorf("Expected node left to be nil, got %v", node.Left)
		}
		if node.Right != nil {
			t.Errorf("Expected node right to be nil, got %v", node.Right)
		}
		if node.height != 1 {
			t.Errorf("Expected node height to be 1, got %v", node.height)
		}
	})

	t.Run("Add node to tree", func(t *testing.T) {
		tree := NewAVLTree[int]()
		tree.AddNode(6)
		tree.AddNode(3)
		tree.AddNode(8)

		root := tree.Root

		if root.Data != 6 {
			t.Errorf("Expected root to hold 6, got %v", root.Data)
		}
		if root.Left.Data != 3 {
			t.Errorf("Expected left of root to hold 3, got %v", root.Left.Data)
		}
		if root.Right.Data != 8 {
			t.Errorf("Expected right of root to hold 8, got %v", root.Right.Data)
		}
	})

	t.Run("Print tree", func(t *testing.T) {
		type testPrintTreeStruct struct {
			args     []int
			expected []int
		}

		var printTreeTests = []testPrintTreeStruct{
			{[]int{}, []int{}},
			{[]int{6, 3, 8, 2, 5, 7, 4}, []int{2, 3, 4, 5, 6, 7, 8}},
			{[]int{15, 25, 35, 45, 55, 20}, []int{15, 20, 25, 35, 45, 55}},
		}

		for _, test := range printTreeTests {
			// create new tree for each test case instance
			tree := NewAVLTree[int]()

			// populate the tree with nodes
			for _, arg := range test.args {
				tree.AddNode(arg)
			}

			printed := tree.Root.PrintTree()

			for i, node := range printed {
				if node != test.expected[i] {
					t.Errorf("Expected %v, got %v", test.expected, printed)
				}
			}

		}
	})

	t.Run("Balances tree", func(t *testing.T) {
		type testBalanceTreeStruct struct {
			args     []int
			expected int
		}

		var balanceTreeTests = []testBalanceTreeStruct{
			{[]int{4, 3, 1}, 3}, // left left
			{[]int{4, 5, 6}, 5}, // right right
			{[]int{4, 2, 3}, 3}, // left right
			{[]int{6, 8, 7}, 7}, // right left
		}

		for _, test := range balanceTreeTests {
			// create new tree for each test case instance
			tree := NewAVLTree[int]()

			// populate the tree with nodes
			for _, arg := range test.args {
				tree.AddNode(arg)
			}

			root := tree.Root
			if root.Data != test.expected {
				t.Errorf("Expected root to hold %v, got %v", test.expected, root.Data)
			}
		}
	})
}
