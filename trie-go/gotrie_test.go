package gotrie

import (
	"testing"
)

func addWordNTimes(word string, n int) (*Node, uint16) {
	trie := New()
	var num uint16

	for i := 0; i < n; i++ {
		num = trie.Add(word)
	}
	return trie, num
}

func TestTrieAdd__HappyPath(t *testing.T) {
	_, num := addWordNTimes("cat", 2)

	if num != 2 {
		t.Errorf("Word count should be 2, got %d", num)
	}
}

func TestTrieGet__HappyPath(t *testing.T) {
	trie, _ := addWordNTimes("cat", 2)
	num := trie.Get("cat")

	if num != 2 {
		t.Errorf("Word count should be 2, got %d", num)
	}
}

func TestTrieGet__NotAdded(t *testing.T) {
	trie, _ := addWordNTimes("cat", 2)
	num := trie.Get("dog")

	if num != 0 {
		t.Errorf("Word count should be 0, got %d", num)
	}
}

func TestTrieDelete__HappyPath(t *testing.T) {
	trie, _ := addWordNTimes("cat", 2)
	num := trie.Delete("cat")

	if num != 1 {
		t.Errorf("Word count should be 1, got %d", num)
	}
}

func TestTrieDelete__RemoveToParentNode(t *testing.T) {
	trie, _ := addWordNTimes("cat", 1)
	trie.Delete("cat")

	numChildren := len(trie.children)
	if numChildren != 0 {
		t.Errorf("Trie should have no children, has %d", numChildren)
	}
}

func TestNodeCount(t *testing.T) {
	trie := New()
	node := newNode('a', trie)
	node.count = node.count + 1
	if node.count != 1 {
		t.Errorf("Node should have count 1, has %d", node.count)
	}
}
