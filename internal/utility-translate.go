package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/null93/superperm/sdk/alpha"
	"github.com/spf13/cobra"
)

var utilityTranslateCmd = &cobra.Command{
	Use:     "translate -a ALPHABET -f FILE",
	Short:   "translate solution from one alphabet to another",
	Example: "translate ./solutions/rotate/4-33.txt --size 4 --from digit --to latin",
	Aliases: []string{"tran", "trans"},
	Args:    cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		n, _ := cmd.Flags().GetInt("size")
		from, _ := cmd.Flags().GetString("from")
		to, _ := cmd.Flags().GetString("to")
		solutionPath := args[0]
		if n < 1 {
			return fmt.Errorf("alphabet size must be at least 1")
		}
		if !alpha.IsSupportedAlphabet(from) {
			return fmt.Errorf("from alphabet not supported")
		}
		if !alpha.IsSupportedAlphabet(to) {
			return fmt.Errorf("to alphabet not supported")
		}
		if _, err := os.Stat(solutionPath); err != nil {
			return fmt.Errorf("file does not exist")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		n, _ := cmd.Flags().GetInt("size")
		from, _ := cmd.Flags().GetString("from")
		to, _ := cmd.Flags().GetString("to")
		solutionPath := args[0]
		solution, _ := os.ReadFile(solutionPath)
		if translator, err := alpha.NewTranslator(from, to, n); err == nil {
			solutionTrimmed := strings.TrimSpace(string(solution))
			if translation, err := translator.Translate(solutionTrimmed); err == nil {
				fmt.Println(translation)
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	},
}

func init() {
	utilityCmd.AddCommand(utilityTranslateCmd)
	utilityTranslateCmd.Flags().SortFlags = true
	utilityTranslateCmd.Flags().IntP("size", "n", 0, "alphabet size")
	utilityTranslateCmd.Flags().StringP("from", "f", "", "source alphabet name")
	utilityTranslateCmd.Flags().StringP("to", "t", "", "target alphabet name")
}
