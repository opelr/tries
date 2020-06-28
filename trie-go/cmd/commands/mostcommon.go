package commands

import (
	"github.com/spf13/cobra"

	"github.com/opelr/trie/gotrie"
)

func init() {
	rootCmd.AddCommand(mostcommonCmd)

	mostcommonCmd.Flags().IntP("number", "n", 10, "Number of words to return.")
}

// mostcommonCmd represents the mostcommon command
var mostcommonCmd = &cobra.Command{
	Use:   "mostcommon",
	Args:  cobra.ExactArgs(1),
	Short: "Find the most common words in a file and their counts.",
	RunE:  mostCommon,
}

func mostCommon(cmd *cobra.Command, args []string) error {
	wordArray, err := readFile(cmd, args)
	if err != nil {
		return err
	}

	// Build Trie
	trie := gotrie.New()
	for _, word := range wordArray {
		trie.Add(word)
	}

	num, err := cmd.Flags().GetInt("number")
	if err != nil {
		return err
	}
	err = trie.MostCommon(num)
	if err != nil {
		return err
	}
	return nil
}
