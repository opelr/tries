package gotrie

import (
	"sort"
)

// Word is a struct for counting the number of times a word occurs in the
// trie.
type Word struct {
	Value string
	Count uint16
}

// Walk returns all words from the trie and their counts.
// Returns an array of Words, one per unique word.
func (trie *Node) Walk() []Word {
	ch := make(chan Word)
	go func() {
		walkTrie(trie, ch)
		close(ch)
	}()

	var words []Word
	for word := range ch {
		words = append(words, word)
	}
	return words
}

// Search finds all words in the trie that match a prefix substring.
func (trie *Node) Search(substring string) ([]Word, error) {
	node, err := trie.Get(substring)

	if node == nil || err != nil {
		return nil, err
	}

	words := node.Walk()
	return words, nil
}

// SortAlpha sorts a slice of Word structs alphabetically by Value.
func SortAlpha(words []Word) []Word {
	sort.Slice(words, func(i, j int) bool {
		return words[i].Value < words[j].Value
	})
	return words
}

// SortAlphaNumeric sorts a slice of Word structs numerically by Count, then
// alphabetically by Value.
func SortAlphaNumeric(words []Word) []Word {
	sort.Slice(words, func(i, j int) bool {
		if words[i].Count > words[j].Count {
			return true
		}
		if words[i].Count < words[j].Count {
			return false
		}
		return words[i].Value < words[j].Value
	})
	return words
}

// getWord is a helper function that returns a Word given a Node.
func (trie *Node) getWord() Word {
	word := ""
	count := trie.count
	for {
		if trie.parent == nil {
			return Word{Value: word, Count: count}
		}
		word = string(trie.value) + word
		trie = trie.parent
	}

}

// walk is a helper function that walks the trie and passes words (via Word
// structs) to a channel.
func walkTrie(trie *Node, ch chan Word) {
	if trie == nil {
		return
	}
	if trie.count > 0 {
		ch <- trie.getWord()
	}
	for _, child := range trie.children {
		walkTrie(child, ch)
	}
}
