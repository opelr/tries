// Package gotrie implements a prefix tree (trie) data structure and methods
// for retrieving information from the tree.
package gotrie

import (
	"fmt"
)

// The Node struct represents a single node in the prefix tree (i.e. trie).
// A Node with no parent is considered the root of the trie, and thus contains
// no value.
type Node struct {
	value    rune
	count    uint16
	parent   *Node
	children map[rune]*Node
}

// New creates a new trie.
func New() *Node {
	return new(Node)
}

// newNode is a helper function for creating a new node in the trie.
// It assigns a value to the node and adds a pointer to the parent node, which
// may be the trie root.
func newNode(value rune, parent *Node) *Node {
	node := New()
	node.value = value
	node.parent = parent
	return node
}

// Add inserts a word into the trie and increments the count.
// Returns a pointer to the Node that was added or updated.
func (trie *Node) Add(word string) (*Node, error) {
	node := trie
	for _, char := range word {
		if node.children == nil {
			node.children = make(map[rune]*Node)
		}
		child, _ := node.children[char]
		if child == nil {
			child = newNode(char, node)
			node.children[char] = child
		}
		node = child
	}
	node.count = node.count + 1
	return node, nil
}

// Get finds a word in the trie and returns the number of times it's appeared
// in the trie.
// Returns a pointer to the Node that was found or nil.
func (trie *Node) Get(word string) (*Node, error) {
	node := trie
	for _, char := range word {
		node, _ = node.children[char]
		if node == nil {
			return nil, fmt.Errorf("Word '%v' doesn't exist in the trie", word)
		}
	}
	return node, nil
}

// Delete removes a word from the trie and cleans up any nodes without a
// reference.
// Returns a pointer to the Node that was updated or nil.
func (trie *Node) Delete(word string) (*Node, error) {
	node, err := trie.Get(word)
	if err != nil {
		return nil, err
	}

	// Decrement the word count.
	if node.count > 0 {
		node.count = node.count - 1
	}

	if !node.isTerminal() {
		return node, nil
	}

	// Clean up nodes without a reference.
	for {
		if !node.isTerminal() || node.isRoot() {
			break
		}
		parent := node.parent
		delete(parent.children, node.value)
		node = parent
	}
	return nil, nil
}

// isTerminal is a helper function that determines if a node is a leaf with no
// count.
func (trie *Node) isTerminal() bool {
	return len(trie.children) == 0 && trie.count <= 0
}

// isRoot is a helper function that determines if a node is the trie root.
func (trie *Node) isRoot() bool {
	return trie.parent == nil && trie.value == 0
}
