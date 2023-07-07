package AVLTree

import (
	"golang.org/x/exp/constraints"
)

type avlTree[T constraints.Ordered] struct {
	Root *node[T]
}

type node[T constraints.Ordered] struct {
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
	return &node[T]{data, nil, nil, 1}
}

// AddNode using the insert function
func (a *avlTree[T]) AddNode(data T) {
	newNode := NewNode(data)
	if a.Root == nil {
		a.Root = newNode
	} else {
		a.Root = insert(a.Root, newNode)
	}
}

// Recursively insert the newnode at the correct place
func insert[T constraints.Ordered](node *node[T], newNode *node[T]) *node[T] {
	if node == nil {
		return newNode
	}

	if newNode.Data < node.Data {
		node.Left = insert(node.Left, newNode)
	} else {
		node.Right = insert(node.Right, newNode)
	}

	// update height of ancestor node
	node.height = maxHeight(getHeight(node.Left), getHeight(node.Right)) + 1

	// get balance factor
	balance := getBalanceFactor(node)

	// if node is imbalanced:
	if balance > 1 || balance < -1 {
		// left left imbalance
		if getBalanceFactor(node) == 2 && getBalanceFactor(node.Left) == 1 {
			return rotateTree[T]("right", node)
		}
		// left right imbalance
		if getBalanceFactor(node) == 2 && getBalanceFactor(node.Left) == -1 {
			node.Left = rotateTree[T]("left", node.Left)
			return rotateTree[T]("right", node)
		}
		// right right imbalance
		if getBalanceFactor(node) == -2 && getBalanceFactor(node.Right) == -1 {
			return rotateTree[T]("left", node)
		}
		// right left imbalance
		if getBalanceFactor(node) == -2 && getBalanceFactor(node.Right) == 1 {
			node.Right = rotateTree[T]("right", node.Right)
			return rotateTree[T]("left", node)
		}
	}

	return node
}

func rotateTree[T constraints.Ordered](direction string, root *node[T]) *node[T] {
	var newRoot *node[T]

	if direction == "left" {
		newRoot = root.Right
		root.Right = newRoot.Left
		newRoot.Left = root
	} else {
		newRoot = root.Left
		root.Left = newRoot.Right
		newRoot.Right = root
	}

	root.height = maxHeight(getHeight(root.Left), getHeight(root.Right)) + 1
	newRoot.height = maxHeight(getHeight(newRoot.Left), getHeight(newRoot.Right)) + 1
	return newRoot
}

func PrintTree[T constraints.Ordered](root *node[T]) []T {
	if root == nil {
		return []T{}
	}

	var result []T
	result = append(result, PrintTree(root.Left)...)
	result = append(result, root.Data)
	result = append(result, PrintTree(root.Right)...)
	return result
}

func getHeight[T constraints.Ordered](root *node[T]) int {
	if root == nil {
		return 0
	}
	return root.height
}

func getBalanceFactor[T constraints.Ordered](root *node[T]) int {
	if root == nil {
		return 0
	}
	return getHeight(root.Left) - getHeight(root.Right)
}

func maxHeight(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
