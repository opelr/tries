package wordio

import (
	"bufio"
	"os"
	"strings"
)

// ReadTxt reads a file at a given path and returns an array of words.
func ReadTxt(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineWords := strings.Fields(line)
		words = append(words, lineWords...)
	}
	return words, scanner.Err()
}

// ToLower takes an array of strings and returns an array where each string is
// cast to lowercase.
func ToLower(stringList []string) []string {
	var words []string
	for _, word := range stringList {
		word = strings.ToLower(word)
		words = append(words, word)
	}
	return words
}

// TrimPunctuation takes an array of strings and returns an array where each
// string is stripped of trailing punctutation.
func TrimPunctuation(stringList []string) []string {
	var words []string
	for _, word := range stringList {
		word = strings.Trim(word, ".,!?;:'\"_-*()[]")
		words = append(words, word)
	}
	return words
}
