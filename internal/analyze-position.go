package internal

import (
	"fmt"
	"os"

	"github.com/null93/superperm/sdk/utils"
	"github.com/spf13/cobra"
)

var analyzePositionCmd = &cobra.Command{
	Use:     "position -a ALPHABET -f FILE",
	Short:   "Histogram of the position of the character in the alphabet",
	Example: "position -a 1234 -f ./solutions/rotate/4-33.txt | less -RS",
	Aliases: []string{"pos"},
	Args:    cobra.NoArgs,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		alphabet, _ := cmd.Flags().GetString("alphabet")
		solutionPath, _ := cmd.Flags().GetString("file")
		if len(alphabet) < 1 {
			return fmt.Errorf("alphabet must be at least 1 character long")
		}
		if _, err := os.Stat(solutionPath); err != nil {
			return fmt.Errorf("file does not exist")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		noColor, _ := cmd.Flags().GetBool("no-color")
		alphabet, _ := cmd.Flags().GetString("alphabet")
		solutionPath, _ := cmd.Flags().GetString("file")
		solution, _ := os.ReadFile(solutionPath)
		utils.DisableColor(noColor)
		utils.PrintHistogram(alphabet, string(solution))
	},
}

func init() {
	analyzeCmd.AddCommand(analyzePositionCmd)
	analyzePositionCmd.Flags().SortFlags = true
	analyzePositionCmd.Flags().Bool("no-color", false, "show colors in output")
	analyzePositionCmd.Flags().StringP("alphabet", "a", "", "exact alphabet used to generate the solution")
	analyzePositionCmd.Flags().StringP("file", "f", "", "path to file with solution")
}
