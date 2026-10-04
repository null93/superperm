package internal

import (
	"github.com/null93/superperm/sdk/utils"
	"github.com/spf13/cobra"
)

var analyzeHeatMapCmd = &cobra.Command{
	Use:     "heatmap FILE",
	Short:   "Histogram of solution with magnitude representing how many permutations it is a part of",
	Example: "heatmap ./solutions/standard/4-33.txt | less -RS",
	Aliases: []string{"heat"},
	Args:    cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		_, _, err := utils.ReadSolution(args[0])
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		noColor, _ := cmd.Flags().GetBool("no-color")
		solution, alphabet, _ := utils.ReadSolution(args[0])
		permutations := utils.Permutations(alphabet)
		utils.DisableColor(noColor)
		utils.PrintHeatMap(permutations, solution)
	},
}

func init() {
	analyzeCmd.AddCommand(analyzeHeatMapCmd)
	analyzeHeatMapCmd.Flags().SortFlags = true
	analyzeHeatMapCmd.Flags().Bool("no-color", false, "show colors in output")
}
