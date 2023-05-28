package internal

import (
	"fmt"
	"os"

	"github.com/null93/superperm/sdk/utils"
	"github.com/spf13/cobra"
)

var analyzeHeatMapCmd = &cobra.Command{
	Use:     "heatmap -a ALPHABET -f FILE",
	Short:   "Histogram of solution with magnitude representing how many permutations it is a part of",
	Example: "heatmap -a 1234 -f ./solutions/rotate/4-33.txt | less -RS",
	Aliases: []string{"heat"},
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
		utils.PrintHeatMap(permutations, string(solution))
	},
}

func init() {
	analyzeCmd.AddCommand(analyzeHeatMapCmd)
	analyzeHeatMapCmd.Flags().SortFlags = true
	analyzeHeatMapCmd.Flags().Bool("no-color", false, "show colors in output")
	analyzeHeatMapCmd.Flags().StringP("alphabet", "a", "", "exact alphabet used to generate the solution")
	analyzeHeatMapCmd.Flags().StringP("file", "f", "", "path to file with solution")
}
