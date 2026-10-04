package internal

import (
	"fmt"
	"os"

	"github.com/null93/superperm/sdk/utils"
	"github.com/spf13/cobra"
)

var analyzeExtractCmd = &cobra.Command{
	Use:     "extract FILE",
	Short:   "Extract permutations found in the solution, aligned to where they occur",
	Example: "extract ./solutions/standard/4-33.txt | less -RS\nextract ./solutions/standard/4-33.txt -s 4-33.svg",
	Aliases: []string{"ext"},
	Args:    cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		_, _, err := utils.ReadSuperpermutation(args[0])
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		noColor, _ := cmd.Flags().GetBool("no-color")
		output, _ := cmd.Flags().GetString("svg-output")
		solution, alphabet, _ := utils.ReadSolution(args[0])
		permutations := utils.Permutations(alphabet)
		utils.DisableColor(noColor)
		if output == "" {
			utils.PrintExtraction(permutations, solution)
			return
		}
		if err := utils.WriteExtraction(permutations, solution, output); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	analyzeCmd.AddCommand(analyzeExtractCmd)
	analyzeExtractCmd.Flags().SortFlags = true
	analyzeExtractCmd.Flags().Bool("no-color", false, "show colors in output")
	analyzeExtractCmd.Flags().StringP("svg-output", "s", "", "write extraction as svg to path instead of printing")
}
