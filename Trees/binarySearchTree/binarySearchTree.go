package binarySearchTree

import "golang.org/x/exp/constraints"

type node[T constraints.Ordered] struct {
	Data  T
	Left  *node[T]
	Right *node[T]
}

type binaryTree[T constraints.Ordered] struct {
	Root *node[T]
}

// NewBinaryTree initialize new binaryTree
func NewBinaryTree[T constraints.Ordered]() *binaryTree[T] {
	binaryTree := new(binaryTree[T])
	binaryTree.Root = nil
	return binaryTree
}

// NewNode Create new node
func NewNode[T constraints.Ordered](data T) *node[T] {
	node := new(node[T])
	node.Data = data
	node.Left = nil
	node.Right = nil
	return node
}

// AddNode Insert
func (b *binaryTree[T]) AddNode(data T) {
	newNode := NewNode(data)
	if b.Root == nil {
		// set Root to new node if empty
		b.Root = newNode
	} else {
		// traverse tree to find correct position for new node
		current := b.Root
		for {
			if (data < current.Data) {
				if current.Left == nil {
					current.Left = newNode
					return
				}
				current = current.Left
			} else {
				if current.Right == nil {
					current.Right = newNode
					return
				}
				current = current.Right
			}
		}
	}
}


// Pre-order traversal while printing values of nodes
func (b *binaryTree[T]) PreOrder(root *node[T]) []T {
	var treeNodes []T
	traverse("pre", root, &treeNodes)
	return treeNodes
}


// In-order traversal while printing values of nodes
func (b *binaryTree[T]) InOrder(root *node[T]) []T {
	var treeNodes []T
	traverse("in", root, &treeNodes)
	return treeNodes
}

// Post-order traversal while printing values of nodes
func (b *binaryTree[T]) PostOrder(root *node[T]) []T {
	var treeNodes []T
	traverse("post", root, &treeNodes)
	return treeNodes
}

func traverse[T constraints.Ordered](order string, root *node[T], treeNodes *[]T) {
	if root == nil {
		return
	}
	switch order {
	case "pre":
		*treeNodes = append(*treeNodes, root.Data)
		traverse("pre", root.Left, treeNodes)
		traverse("pre", root.Right, treeNodes)
	
	case "in":
		traverse("in", root.Left, treeNodes)
		*treeNodes = append(*treeNodes, root.Data)
		traverse("in", root.Right, treeNodes)

	case "post":
		traverse("post", root.Left, treeNodes)
		traverse("post", root.Right, treeNodes)
		*treeNodes = append(*treeNodes, root.Data)
		
	}
}
