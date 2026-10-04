package internal

import (
	"github.com/null93/superperm/sdk/utils"
	"github.com/spf13/cobra"
)

var analyzeExtractCmd = &cobra.Command{
	Use:     "extract FILE",
	Short:   "Extract permutations found in the solution, aligned to where they occur",
	Example: "extract ./solutions/standard/4-33.txt | less -RS",
	Aliases: []string{"ext"},
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
		utils.PrintExtraction(permutations, solution)
	},
}

func init() {
	analyzeCmd.AddCommand(analyzeExtractCmd)
	analyzeExtractCmd.Flags().SortFlags = true
	analyzeExtractCmd.Flags().Bool("no-color", false, "show colors in output")
}
