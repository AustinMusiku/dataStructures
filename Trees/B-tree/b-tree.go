package B_tree

import (
	"bytes"
	"fmt"
	"math"
	"sort"
)

type BTree struct {
	order int
	root  *node
}

type node struct {
	order    int
	parent   *node
	inodes   Inodes
	children Pnodes
}

type inode []byte

type Inodes []inode
type Pnodes []*node

func New(order int) (*BTree, error) {
	if order < 2 {
		return nil, fmt.Errorf("order must be greater than 1")
	}

	bTree := &BTree{
		order: order,
		root:  nil,
	}

	return bTree, nil
}

func NewNode(order int, parent *node) (*node, error) {
	if order < 2 {
		return nil, fmt.Errorf("order must be greater than 1")
	}

	node := &node{
		order:    order,
		parent:   parent,
		inodes:   make([]inode, 0, order-1),
		children: make([]*node, 0, order),
	}

	return node, nil
}

func (b *BTree) Insert(key []byte) error {
	// empty value
	if len(key) == 0 {
		return fmt.Errorf("tree insert: empty key")
	}

	// empty root
	if b.root == nil {
		node, err := NewNode(b.order, nil)
		if err != nil {
			return fmt.Errorf("tree insert: %v", err)
		}

		b.root = node
	}

	var err error
	b.root, err = b.root.put(key)

	return err
}

func (n *node) put(key []byte) (*node, error) {
	idx := sort.Search(len(n.inodes), func(i int) bool {
		return bytes.Compare(n.inodes[i], key) >= 0
	})

	if idx < len(n.inodes) && bytes.Equal(key, n.inodes[idx]) {
		return nil, fmt.Errorf("node put: key already exists")
	}

	// if internal node, recurse to child
	if len(n.children) > 0 {
		child := n.children[idx]
		child.put(key)
		if len(n.inodes) > n.order-1 {
			return n.rebalance(n.inodes)
		}
		return n, nil
	}

	// if leaf node is not full, insert
	if len(n.inodes) < n.order-1 {
		n.inodes.InsertAt(idx, key)
		return n, nil
	}

	// if leaf node is full, split
	keys := stretchInodes(n, key)
	return n.rebalance(keys)
}

func (n *node) rebalance(keys Inodes) (*node, error) {
	right, midPoint := splitKeys(n, keys)
	parent, inodeInsertionIdx := n.parent, 0

	if parent != nil {
		inodeInsertionIdx = n.parent.inodes.getInsertionIdx(midPoint)
	} else { // if root node is full, create new root
		parent, _ := NewNode(n.order, nil)
		parent.children = append(parent.children, n)
		n.parent = parent
		right.parent = parent
	}

	n.parent.inodes.InsertAt(inodeInsertionIdx, midPoint)
	n.parent.children.InsertAt(inodeInsertionIdx+1, right)

	return n.parent, nil
}

func stretchInodes(n *node, key inode) Inodes {
	var keys Inodes = make([]inode, (n.order-1)*2)
	copy(keys, n.inodes)
	keys = append(keys, nil)
	keys.InsertAt(n.inodes.getInsertionIdx(key), key)

	return keys[:len(n.inodes)+1]
}

func splitKeys(n *node, keys Inodes) (*node, inode) {
	child2, _ := NewNode(n.order, n.parent)

	splitIdx := int(math.Floor(float64(len(keys)) / 2))
	n.inodes = n.inodes[:0]
	n.inodes = keys[:splitIdx]
	child2.inodes = keys[splitIdx+1:]

	if len(n.children) > 0 {
		child2.children = n.children[splitIdx+1:]
		n.children = n.children[:splitIdx+1]
	}

	midPoint := keys[splitIdx]

	return child2, midPoint
}

func (i *Inodes) getInsertionIdx(key []byte) int {
	return sort.Search(len(*i), func(idx int) bool {
		return bytes.Compare((*i)[idx], key) >= 0
	})
}

func (i *Inodes) InsertAt(idx int, value []byte) {
	*i = append(*i, nil)
	copy((*i)[idx+1:], (*i)[idx:])
	(*i)[idx] = value
}

func (i *Inodes) RemoveAt(idx int) {
	copy((*i)[:idx], (*i)[idx+1:])
	*i = (*i)[:len(*i)-1]
}

func (i *Pnodes) InsertAt(idx int, value *node) {
	*i = append(*i, nil)
	copy((*i)[idx+1:], (*i)[idx:])
	(*i)[idx] = value
}

func (i *Pnodes) RemoveAt(idx int) {
	copy((*i)[idx:], (*i)[idx+1:])
	*i = (*i)[:len(*i)-1]
}
