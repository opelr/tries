package commands

import (
	"fmt"
	"os"

	"github.com/opelr/trie/gotrie/internal/wordio"
	"github.com/spf13/cobra"

	"github.com/MakeNowJust/heredoc"
)

func init() {
	rootCmd.PersistentFlags().BoolP("ignore-case", "i", false, "Make commands case insensitive. By default, *gotrie* is case sensitive.")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "gotrie <command> [flags]",
	Short:         "Trie CLI written in Go!",
	SilenceErrors: true,
	SilenceUsage:  true,
	Example: heredoc.Doc(`
	$ gotrie mostcommon <path-to-file.txt> -n 5
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
	filePath := args[0]
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
