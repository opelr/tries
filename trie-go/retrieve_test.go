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
	if wordCounter.Value != "cat" {
		t.Errorf("Value should be 'cat', got %s", wordCounter.Value)
	}
	if wordCounter.Count != 2 {
		t.Errorf("Word count should be 2, got %d", wordCounter.Count)
	}
}

func TestTrieSearch(t *testing.T) {
	trie := New()
	words := []string{"cat", "car", "carp", "carpenter", "dog"}
	for _, word := range words {
		trie.Add(word)
	}

	expectedWords := [4]string{"cat", "car", "carp", "carpenter"}
	expectedWordsMap := make(map[string]bool)
	for _, word := range expectedWords {
		expectedWordsMap[word] = false
	}

	// Internal helper method for determining if a word is contained in
	// `expectedWords`.
	in := func(word string) bool {
		for _, x := range expectedWords {
			if word == x {
				return true
			}
		}
		return false
	}

	foundWords, _ := trie.Search("ca")

	for _, word := range foundWords {
		if !in(word.Value) {
			t.Errorf("Word '%v' was returned by Search, and it should not be", word)
		}
		expectedWordsMap[word.Value] = true
	}

	for key, value := range expectedWordsMap {
		if value == false {
			t.Errorf("Word '%v' was not returned by Search, and should have been", key)
		}
	}

}
