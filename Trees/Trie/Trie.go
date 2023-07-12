package Trie

import (
	"sync"
)

// Limit the alphabet to lowercase letters a-z
const AlphabetSize = 26

type node struct {
	mu         sync.Mutex
	children   [AlphabetSize]*node
	isTerminal bool
}

type Trie struct {
	root *node
}

func NewNode() *node {
	return &node{}
}

func NewTrie() *Trie {
	return &Trie{root: NewNode()}
}

// --------------------
// Trie methods
// --------------------
// Insert
func (t *Trie) Insert(word string) bool {
	return t.root.insert(word)
}

// Search
func (t *Trie) Search(word string) bool {
	return t.root.search(word)
}

// Delete
func (t *Trie) Delete(word string) bool {
	return t.root.delete(word)
}

// --------------------
// Node methods
// --------------------
// Recursively insert a word into the trie starting at the root node.
// Returns true if the word is inserted, false otherwise
func (n *node) insert(word string) bool {
	n.mu.Lock()
	defer n.mu.Unlock()

	if len(word) == 0 {
		return false
	}

	current := n

	letterIdx := word[:1][0] - 'a'

	if current.children[letterIdx] == nil {
		current.children[letterIdx] = NewNode()
	}

	current = current.children[letterIdx]

	if len(word) == 1 {
		current.isTerminal = true
		return true
	}

	return current.insert(word[1:])
}

// Recursively search for a word in the trie starting at the root node.
// Returns true if the word is found, false otherwise
func (n *node) search(word string) bool {
	current := n

	if len(word) == 0 {
		if current.isTerminal {
			return true
		} else {
			return false
		}
	}

	letterIdx := word[:1][0] - 'a'

	if current.children[letterIdx] != nil {
		next := current.children[letterIdx]
		return next.search(word[1:])
	}

	return false
}

// Recursively delete a word from the trie starting at the root node.
// Returns true if the word is deleted, false otherwise
func (n *node) delete(word string) bool {
	n.mu.Lock()
	defer n.mu.Unlock()

	current := n

	if len(word) == 0 {
		if current.isTerminal {
			current.isTerminal = false
			return true
		} else {
			return false
		}
	}

	letterIdx := word[:1][0] - 'a'

	if current.children[letterIdx] != nil {
		next := current.children[letterIdx]
		delete := next.delete(word[1:])

		if delete {
			// if the node has no children and is not a terminal node, delete the node
			if !next.isTerminal && next.Children() == 0 {
				current.children[letterIdx] = nil
			}
		}

		return delete

	}

	return false
}

// check if a node has no children
func (n *node) Children() int {
	count := 0
	for _, child := range n.children {
		if child != nil {
			count++
		}
	}

	return count
}
