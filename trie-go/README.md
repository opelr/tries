# Gotrie

## Module

`go get github.com/opelr/trie/gotrie`

```go
package main

import "github.com/opelr/trie/gotrie"

func main() {
    trie := gotrie.New()
    trie.Add("cat")
    trie.Add("dog")
    words, err := trie.Search("ca")
}
```

## Command Line Interface

1. `go build cmd/gotrie.go`
1. `./gotrie --help`
