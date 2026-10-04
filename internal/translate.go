package internal

import (
	"fmt"

	"github.com/null93/superperm/sdk/alpha"
	"github.com/null93/superperm/sdk/utils"
	"github.com/spf13/cobra"
)

var translateCmd = &cobra.Command{
	Use:     "translate FILE ALPHABET",
	Short:   "Translate solution from its alphabet to another",
	Example: "translate ./solutions/standard/4-33.txt 1234\nsuperperm perms ABCD | xargs -n1 superperm translate ./solutions/shortest/4-33.txt",
	Aliases: []string{"tran", "trans"},
	Args:    cobra.ExactArgs(2),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		target := args[1]
		if errAlpha := alpha.Validate(target); errAlpha != nil {
			return errAlpha
		}
		_, alphabet, err := utils.ReadSolution(args[0])
		if err != nil {
			return err
		}
		if len(target) != len(alphabet) {
			return fmt.Errorf("alphabet must be %d characters long to match solution alphabet %s", len(alphabet), alphabet)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		target := args[1]
		solution, alphabet, _ := utils.ReadSolution(args[0])
		translation, err := alpha.Translate(solution, alphabet, target)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(translation)
	},
}

func init() {
	RootCmd.AddCommand(translateCmd)
	translateCmd.Flags().SortFlags = true
}
