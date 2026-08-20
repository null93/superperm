package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/null93/superperm/sdk/utils"
	"github.com/spf13/cobra"
)

var analyzeExtractCmd = &cobra.Command{
	Use:     "extract -a ALPHABET -f FILE",
	Short:   "Extract permutations found in the solution, aligned to where they occur",
	Example: "extract -a ABCD -f ./solutions/rotate/4-33.txt | less -RS",
	Aliases: []string{"ext"},
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
		permutations := utils.Permutations(alphabet)
		utils.DisableColor(noColor)
		utils.PrintExtraction(permutations, strings.TrimSpace(string(solution)))
	},
}

func init() {
	analyzeCmd.AddCommand(analyzeExtractCmd)
	analyzeExtractCmd.Flags().SortFlags = true
	analyzeExtractCmd.Flags().Bool("no-color", false, "show colors in output")
	analyzeExtractCmd.Flags().StringP("alphabet", "a", "", "exact alphabet used to generate the solution")
	analyzeExtractCmd.Flags().StringP("file", "f", "", "path to file with solution")
}
