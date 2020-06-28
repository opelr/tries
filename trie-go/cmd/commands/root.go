package commands

import (
	"fmt"
	"os"

	"github.com/opelr/trie/gotrie"
	"github.com/opelr/trie/gotrie/internal/wordio"
	"github.com/spf13/cobra"

	"github.com/MakeNowJust/heredoc"
)

func init() {
	rootCmd.PersistentFlags().BoolP(
		"ignore-case", "i", false,
		"Make commands case insensitive. By default, gotrie is case sensitive.",
	)
	rootCmd.PersistentFlags().BoolP(
		"no-counts", "x", false,
		`If --no-counts, gotrie will only print words for each command.
		By default, gotrie prints counts and words.`,
	)
	rootCmd.PersistentFlags().BoolP(
		"sort-alpha", "a", false,
		"If passed, gotrie will sort words alphabetically (default: sort numerically, then alphabetically).",
	)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "gotrie <command> [flags]",
	Short:         "Trie CLI written in Go!",
	SilenceErrors: true,
	SilenceUsage:  true,
	Example: heredoc.Doc(`
	$ gotrie list <path-to-file.txt> -i
	$ gotrie search "In" <path-to-file.txt> -x
	`),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readFile(cmd *cobra.Command, args []string) ([]string, error) {
	filePath := args[len(args)-1]
	wordArray, err := wordio.ReadTxt(filePath)
	if err != nil {
		return nil, err
	}

	ignoreCase, err := cmd.Flags().GetBool("ignore-case")
	if err != nil {
		return nil, err
	}
	if ignoreCase {
		wordArray = wordio.ToLower(wordArray)
	}

	wordArray = wordio.TrimPunctuation(wordArray)
	return wordArray, nil
}

func buildTrie(cmd *cobra.Command, args []string) (*gotrie.Node, error) {
	wordArray, err := readFile(cmd, args)
	if err != nil {
		return nil, err
	}

	// Build Trie
	trie := gotrie.New()
	for _, word := range wordArray {
		trie.Add(word)
	}
	return trie, nil
}

func printWords(cmd *cobra.Command, args []string, words []gotrie.Word) error {
	// Sort
	sortAlpha, err := cmd.Flags().GetBool("sort-alpha")
	if err != nil {
		return err
	}

	if sortAlpha {
		words = gotrie.SortAlpha(words)
	} else {
		words = gotrie.SortAlphaNumeric(words)
	}

	// Print
	noCounts, err := cmd.Flags().GetBool("no-counts")
	if err != nil {
		return err
	}
	for _, word := range words {
		if noCounts {
			fmt.Println("  ", word.Value)
		} else {
			fmt.Println("  ", word.Value, word.Count)
		}
	}

	return nil
}
