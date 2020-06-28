package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Args:  cobra.ExactArgs(1),
	Short: "Return all unique words in a file and their counts.",
	RunE:  list,
}

func list(cmd *cobra.Command, args []string) error {
	trie, err := buildTrie(cmd, args)
	if err != nil {
		return err
	}

	words := trie.Walk()

	// Print to CLI
	err = printWords(cmd, args, words)
	return nil
}
