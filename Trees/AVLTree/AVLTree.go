package AVLTree

import (
	"sync"

	"golang.org/x/exp/constraints"
)

type avlTree[T constraints.Ordered] struct {
	Root *node[T]
}

type node[T constraints.Ordered] struct {
	mu     sync.Mutex
	Data   T
	Left   *node[T]
	Right  *node[T]
	height int
}

func NewAVLTree[T constraints.Ordered]() *avlTree[T] {
	newAVLTree := new(avlTree[T])
	newAVLTree.Root = nil
	return newAVLTree
}

func NewNode[T constraints.Ordered](data T) *node[T] {
	return &node[T]{
		Data:   data,
		Left:   nil,
		Right:  nil,
		height: 1,
	}
}

// AddNode using the insert function
func (a *avlTree[T]) AddNode(data T) {
	newNode := NewNode(data)
	if a.Root == nil {
		a.Root = newNode
	} else {
		a.Root = a.Root.insert(newNode)
	}
}

// Recursively insert the newnode at the correct place
func (n *node[T]) insert(newNode *node[T]) *node[T] {
	if n == nil {
		return newNode
	}

	if newNode.Data < n.Data {
		n.Left = n.Left.insert(newNode)
	} else {
		n.Right = n.Right.insert(newNode)
	}

	// update height of ancestor node
	n.height = maxHeight(n.Left.getHeight(), n.Right.getHeight()) + 1

	// get balance factor
	balance := n.getBalanceFactor()

	// if node is imbalanced:
	if balance > 1 || balance < -1 {
		// left left imbalance
		if n.getBalanceFactor() == 2 && n.Left.getBalanceFactor() == 1 {
			return n.rotateTree("right")
		}
		// left right imbalance
		if n.getBalanceFactor() == 2 && n.Left.getBalanceFactor() == -1 {
			n.Left = n.Left.rotateTree("left")
			return n.rotateTree("right")
		}
		// right right imbalance
		if n.getBalanceFactor() == -2 && n.Right.getBalanceFactor() == -1 {
			return n.rotateTree("left")
		}
		// right left imbalance
		if n.getBalanceFactor() == -2 && n.Right.getBalanceFactor() == 1 {
			n.Right = n.Right.rotateTree("right")
			return n.rotateTree("left")
		}
	}

	return n
}

// Balance the tree by rotating the subtree rooted at n in the given direction
func (n *node[T]) rotateTree(direction string) *node[T] {
	n.mu.Lock()
	defer n.mu.Unlock()

	var newRoot *node[T]

	if direction == "left" {
		newRoot = n.Right
		n.Right = newRoot.Left
		newRoot.Left = n
	} else {
		newRoot = n.Left
		n.Left = newRoot.Right
		newRoot.Right = n
	}

	n.height = maxHeight(n.Left.getHeight(), n.Right.getHeight()) + 1
	newRoot.height = maxHeight(n.Left.getHeight(), n.Right.getHeight()) + 1
	return newRoot
}

// Returns the items slice of the subtree rooted at n
func (n *node[T]) PrintTree() []T {
	if n == nil {
		return []T{}
	}

	var result []T
	result = append(result, n.Left.PrintTree()...)
	result = append(result, n.Data)
	result = append(result, n.Right.PrintTree()...)
	return result
}

func (n *node[T]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *node[T]) getBalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.Left.getHeight() - n.Right.getHeight()
}

func maxHeight(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
