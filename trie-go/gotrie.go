package gotrie

import (
	"fmt"
	"math"
	"sort"
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
func (trie *Node) Add(word string) uint16 {
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
	return node.count
}

// Get finds a word in the trie and returns the number of times it's appeared
// in the trie.
func (trie *Node) Get(word string) uint16 {
	node := trie
	for _, char := range word {
		node, _ = node.children[char]
		if node == nil {
			return 0
		}
	}
	return node.count
}

// Delete removes a word from the trie and cleans up any nodes without a
// reference.
func (trie *Node) Delete(word string) uint16 {
	node := trie
	for _, char := range word {
		child, _ := node.children[char]
		if child == nil {
			return 0
		}
		node = child
	}

	// Decrement the word count.
	if node.count > 0 {
		node.count = node.count - 1
	}

	if !node.isTerminal() {
		return node.count
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
	return 0
}

// func (trie *Node) Search(substring string) []string {
// 	var words make([]string)
// 	return words
// }

// isTerminal is a helper function that determines if a node is a leaf with no
// count.
func (trie *Node) isTerminal() bool {
	return len(trie.children) == 0 && trie.count <= 0
}

// isRoot is a helper function that determines if a node is the trie root.
func (trie *Node) isRoot() bool {
	return trie.parent == nil && trie.value == 0
}

// MostCommon prints the _n_ most common words in the trie, in descending order.
func (trie *Node) MostCommon(num int) error {
	if num <= 0 {
		return fmt.Errorf("Number must be greater than zero: %d", num)
	}

	words := trie.Walk()

	// Sort in descending order by number of appearances, then an alphabetical
	// order.
	sort.Slice(words, func(i, j int) bool {
		if words[i].count > words[j].count {
			return true
		}
		if words[i].count < words[j].count {
			return false
		}
		return words[i].value < words[j].value
	})
	numWords := int(math.Min(float64(len(words)), float64(num)))

	fmt.Printf("The %d most-common words:\n", numWords)
	for _, count := range words[:numWords] {
		fmt.Println("  ", count.value, count.count)
	}
	return nil
}
