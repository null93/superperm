package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/null93/superperm/sdk/utils"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:     "verify FILE",
	Short:   "Verify that a solution contains every permutation of its alphabet",
	Example: "verify ./solutions/shortest/4-33.txt",
	Aliases: []string{"check"},
	Args:    cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		_, _, err := utils.ReadSolution(args[0])
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		solution, alphabet, _ := utils.ReadSolution(args[0])
		n := len(alphabet)
		missing := utils.MissingPermutations(alphabet, solution)
		if len(missing) > 0 {
			fmt.Printf("invalid superpermutation for n=%d with alphabet %s and length %d, missing %d of %d permutations:\n", n, alphabet, len(solution), len(missing), utils.Factorial(n))
			fmt.Println(strings.Join(missing, "\n"))
			os.Exit(1)
		}
		fmt.Printf("valid superpermutation for n=%d with alphabet %s and length %d\n", n, alphabet, len(solution))
	},
}

func init() {
	RootCmd.AddCommand(verifyCmd)
	verifyCmd.Flags().SortFlags = true
}
