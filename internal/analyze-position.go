package internal

import (
	"github.com/null93/superperm/sdk/utils"
	"github.com/spf13/cobra"
)

var analyzePositionCmd = &cobra.Command{
	Use:     "position FILE",
	Short:   "Histogram of the position of the character in the alphabet",
	Example: "position ./solutions/standard/4-33.txt | less -RS",
	Aliases: []string{"pos"},
	Args:    cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		_, _, err := utils.ReadSolution(args[0])
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		noColor, _ := cmd.Flags().GetBool("no-color")
		solution, alphabet, _ := utils.ReadSolution(args[0])
		utils.DisableColor(noColor)
		utils.PrintHistogram(alphabet, solution)
	},
}

func init() {
	analyzeCmd.AddCommand(analyzePositionCmd)
	analyzePositionCmd.Flags().SortFlags = true
	analyzePositionCmd.Flags().Bool("no-color", false, "show colors in output")
}
