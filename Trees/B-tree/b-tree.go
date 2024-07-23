package B_tree

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strings"
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

func (b *BTree) Search(key []byte) ([]byte, bool) {
	if len(key) == 0 {
		return nil, false
	}
	if b.root == nil {
		return nil, false
	}

	return b.root.search(key)
}

func (b *BTree) Delete(key []byte) error {
	if len(key) == 0 {
		return fmt.Errorf("tree delete: empty key")
	}

	if b.root == nil {
		return fmt.Errorf("tree delete: empty tree")
	}

	b.root.delete(key)
	return nil
}

func (b *BTree) Visualize() string {
	return b.root.Visualize()
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

func (n *node) search(key []byte) ([]byte, bool) {
	idx := sort.Search(len(n.inodes), func(i int) bool {
		return bytes.Compare(n.inodes[i], key) >= 0
	})

	if idx < len(n.inodes) && bytes.Equal(key, n.inodes[idx]) {
		return n.inodes[idx], true
	}

	if len(n.children) == 0 {
		return nil, false
	}

	return n.children[idx].search(key)
}

func (n *node) delete(key []byte) error {
	idx := sort.Search(len(n.inodes), func(i int) bool {
		return bytes.Compare(n.inodes[i], key) >= 0
	})

	if idx < len(n.inodes) && bytes.Equal(n.inodes[idx], key) {
		err := delete(n, idx)
		return err
	}

	if n.children[idx] != nil {
		n.children[idx].delete(key)
		// TODO: recursively borrow an key(aka inode) from L/R sibling or merge with either L/R sibling
		if float64(len(n.children[idx].inodes)) < math.Floor((float64(n.order-1))/2) && n.parent != nil {
			// borrow from left sibling
			if idx > 0 && float64(len(n.children[idx-1].inodes)) > math.Floor((float64(n.order-1))/2) {
				// move key from left child to parent
				fromParent := n.inodes[idx-1]
				n.inodes[idx-1] = n.children[idx-1].inodes[len(n.children[idx-1].inodes)-1]
				n.children[idx].inodes.InsertAt(0, fromParent)
				n.children[idx-1].inodes.RemoveAt(len(n.children[idx-1].inodes) - 1)

				// move child node
				n.children[idx].children.InsertAt(0, n.children[idx-1].children[len(n.children[idx-1].children)-1])
				n.children[idx-1].children.RemoveAt(len(n.children[idx-1].children) - 1)
			} else if idx < len(n.children)-1 && float64(len(n.children[idx+1].inodes)) > math.Floor((float64(n.order-1))/2) {
				// move key from right child to parent
				fromParent := n.inodes[idx]
				n.inodes[idx] = n.children[idx+1].inodes[0]
				n.children[idx].inodes = append(n.children[idx].inodes, fromParent)
				n.children[idx+1].inodes.RemoveAt(0)
				// move child node
				n.children[idx].children = append(n.children[idx].children, n.children[idx+1].children[0])
				n.children[idx+1].children.RemoveAt(0)
			} else {
				// merge into the left sibling
				if idx > 0 {
					n.children[idx-1].inodes = append(n.children[idx-1].inodes, n.inodes[idx-1])
					n.children[idx-1].inodes = append(n.children[idx-1].inodes, n.children[idx].inodes...)
					n.children[idx-1].children = append(n.children[idx-1].children, n.children[idx].children...)
					n.inodes.RemoveAt(idx - 1)
					n.children.RemoveAt(idx)
				} else { // merge into the right sibling
					n.children[idx].inodes = append(n.children[idx].inodes, n.inodes[idx])
					n.children[idx].inodes = append(n.children[idx].inodes, n.children[idx+1].inodes...)
					n.children[idx].children = append(n.children[idx].children, n.children[idx+1].children...)
					n.inodes.RemoveAt(idx)
					n.children.RemoveAt(idx + 1)
				}
			}
		}
	} else {
		return fmt.Errorf("node delete: key not found")
	}

	return nil
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

func delete(n *node, deletionIdx int) error {
	if len(n.children) > 0 { // internal node
		// replace with inorder predecessor
		successor := n.children[deletionIdx].getPredecessor()
		// replace with inorder successor
		if successor == nil {
			successor = n.children[deletionIdx].getSuccessor()
		}

		if successor != nil {
			n.inodes[deletionIdx] = successor
			return nil
		}

		// merge left & right children
		leftChild := n.children[deletionIdx]
		rightChild := n.children[deletionIdx+1]

		leftChild.inodes = append(leftChild.inodes, rightChild.inodes...)
		leftChild.children = append(leftChild.children, rightChild.children...)

		n.inodes.RemoveAt(deletionIdx)
		n.children.RemoveAt(deletionIdx + 1)

		return nil
	}

	// otherwise treat this(n) as a leaf node henceforth
	if len(n.inodes) > ((n.order-1)/2)+1 {
		n.inodes.RemoveAt(deletionIdx)
		return nil
	}

	var donation inode
	var leftSibling, rightSibling *node
	parentDeletionIdx := sort.Search(len(n.parent.inodes), func(i int) bool {
		return bytes.Compare(n.parent.inodes[i], n.inodes[deletionIdx]) >= 0
	})

	// borrow from left sibling
	if parentDeletionIdx > 0 {
		leftSibling = n.parent.children[parentDeletionIdx-1]
		if len(leftSibling.inodes) > (n.order-1)/2 {
			donation = leftSibling.inodes[len(leftSibling.inodes)-1]
			leftSibling.inodes.RemoveAt(len(leftSibling.inodes) - 1)
		}
	}

	// if left has none to donate, borrow from right sibling
	if donation == nil && parentDeletionIdx < len(n.parent.children)-1 {
		rightSibling = n.parent.children[parentDeletionIdx+1]
		if len(rightSibling.inodes) > (n.order-1)/2 {
			donation = rightSibling.inodes[0]
			leftSibling.inodes.RemoveAt(0)
		}
	}

	// if there was a donation, go ahead and accept it
	if donation != nil {
		fromParent := n.parent.inodes[parentDeletionIdx]
		n.parent.inodes[parentDeletionIdx] = donation
		n.inodes.InsertAt(n.inodes.getInsertionIdx(fromParent), fromParent)

		return nil
	}

	// merge with left/right sibling
	n.inodes.RemoveAt(deletionIdx)
	// if leftSibling != nil {
	// 	for i := len(leftSibling.inodes) - 1; i >= 0; i-- {
	// 		n.inodes.InsertAt(0, leftSibling.inodes[i])
	// 	}
	// 	n.parent.children.RemoveAt(parentDeletionIdx - 1)
	// } else {
	// 	n.inodes = append(n.inodes, rightSibling.inodes...)
	// 	n.parent.children.RemoveAt(parentDeletionIdx + 1)
	// }

	return nil
}

func (n *node) Visualize() string {
	if n == nil {
		return ""
	}

	var result strings.Builder
	q := []*node{n}
	for len(q) > 0 {
		var nextQ []*node
		for _, node := range q {
			var localResult string
			for i, inode := range node.inodes {
				localResult += string(inode)
				if i < len(node.inodes)-1 {
					localResult += ","
				}
			}
			result.WriteString(fmt.Sprintf("(%v)\t", localResult))
			nextQ = append(nextQ, node.children...)
		}
		result.WriteString("\n")
		q = nextQ
	}

	return result.String()
}

func (n *node) getPredecessor() inode {
	if len(n.children) == 0 {
		if float64(len(n.inodes)) > math.Floor((float64(n.order-1))/2) {
			p := n.inodes[len(n.inodes)-1]
			n.inodes = n.inodes[:len(n.inodes)-1]

			return p
		}

		return nil
	}

	return n.children[len(n.children)-1].getPredecessor()
}

func (n *node) getSuccessor() inode {
	if len(n.children) == 0 {
		if float64(len(n.inodes)) > math.Floor((float64(n.order-1))/2) {
			s := n.inodes[0]
			n.inodes = n.inodes[1:]

			return s
		}

		return nil

	}

	return n.children[0].getSuccessor()
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
