package gotrie

// Counter is a struct for counting the number of times a word occurs in the
// trie.
type Counter struct {
	value string
	count uint16
}

// Walk returns all words from the trie and their counts.
// Returns an array of Counters, one per unique word.
func (trie *Node) Walk() []Counter {
	ch := make(chan Counter)
	go func() {
		walkTrie(trie, ch)
		close(ch)
	}()

	var words []Counter
	for word := range ch {
		words = append(words, word)
	}
	return words
}

// getWord is a helper function that returns a Counter given a Node.
func (trie *Node) getWord() Counter {
	word := ""
	count := trie.count
	for {
		if trie.parent == nil {
			return Counter{value: word, count: count}
		}
		word = string(trie.value) + word
		trie = trie.parent
	}

}

// walk is a helper function that walks the trie and passes words (via Counter
// structs) to a channel.
func walkTrie(trie *Node, ch chan Counter) {
	if trie == nil {
		return
	}
	for _, child := range trie.children {
		if child.count > 0 {
			ch <- child.getWord()
		}
		walkTrie(child, ch)
	}
}
