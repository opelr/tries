package gotrie

import (
	"testing"
)

func TestGetWord(t *testing.T) {
	trie := New()
	trie.Add("cat")
	trie.Add("cat")

	node := trie.children['c'].children['a'].children['t']
	if node == nil {
		t.Errorf("Node is nil, should be type *Node")
	}
	wordCounter := node.getWord()
	if wordCounter.value != "cat" {
		t.Errorf("Value should be 'cat', got %s", wordCounter.value)
	}
	if wordCounter.count != 2 {
		t.Errorf("Word count should be 2, got %d", wordCounter.count)
	}
}
