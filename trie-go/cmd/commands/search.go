package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd)
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Args:  cobra.ExactArgs(2),
	Short: "Finds all words that match a prefix substring.",
	RunE:  search,
}

func search(cmd *cobra.Command, args []string) error {
	trie, err := buildTrie(cmd, args)
	if err != nil {
		return err
	}

	substring := args[0]
	words, err := trie.Search(substring)
	if err != nil {
		return err
	}

	// Print to CLI
	err = printWords(cmd, args, words)
	return nil
}
