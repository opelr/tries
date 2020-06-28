package gotrie

import (
	"testing"
)

func addWordNTimes(word string, n int) *Node {
	trie := New()

	for i := 0; i < n; i++ {
		trie.Add(word)
	}
	return trie
}

func TestTrieAdd__HappyPath(t *testing.T) {
	trie := addWordNTimes("cat", 2)
	node := trie.children['c'].children['a'].children['t']

	if node.count != 2 {
		t.Errorf("Word count should be 2, got %d", node.count)
	}
}

func TestTrieGet__HappyPath(t *testing.T) {
	trie := addWordNTimes("cat", 2)
	node, _ := trie.Get("cat")

	if node.count != 2 {
		t.Errorf("Word count should be 2, got %d", node.count)
	}
}

func TestTrieGet__NotAdded(t *testing.T) {
	trie := New()
	trie.Add("cat")
	trie.Add("cat")
	node, err := trie.Get("dog")

	if node != nil {
		t.Errorf("Get should return nil when a word doesn't exist in the trie, returned %v", node)
	}
	if err == nil {
		t.Errorf("Get should return an error when a word doesn't exist in the trie.")
	}
}

func TestTrieDelete__HappyPath(t *testing.T) {
	trie := addWordNTimes("cat", 2)
	node, _ := trie.Delete("cat")

	if node.count != 1 {
		t.Errorf("Word count should be 1, got %d", node.count)
	}
}

func TestTrieDelete__RemoveToParentNode(t *testing.T) {
	trie := addWordNTimes("cat", 1)
	trie.Delete("cat")

	numChildren := len(trie.children)
	if numChildren != 0 {
		t.Errorf("Trie should have no children, has %d", numChildren)
	}
}
