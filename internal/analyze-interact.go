package internal

import (
	"fmt"

	"github.com/null93/superperm/sdk/interact"
	"github.com/null93/superperm/sdk/utils"
	"github.com/spf13/cobra"
)

var analyzeInteractCmd = &cobra.Command{
	Use:     "interact FILE",
	Short:   "Rotate cycles of a solution in place and watch the layout change",
	Example: "interact ./solutions/shortest/4-33.txt",
	Aliases: []string{"inter"},
	Args:    cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		_, _, err := utils.ReadSolution(args[0])
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		solution, alphabet, _ := utils.ReadSolution(args[0])
		permutations := utils.Permutations(alphabet)
		if err := interact.Run(permutations, solution); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	analyzeCmd.AddCommand(analyzeInteractCmd)
	analyzeInteractCmd.Flags().SortFlags = true
}
