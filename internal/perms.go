package internal

import (
	"fmt"
	"strings"

	"github.com/null93/superperm/sdk/alpha"
	"github.com/null93/superperm/sdk/utils"
	"github.com/spf13/cobra"
)

var permsCmd = &cobra.Command{
	Use:     "perms ALPHABET",
	Short:   "List all permutations of an alphabet, order matters",
	Example: "perms ABCD",
	Aliases: []string{"perm"},
	Args:    cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return alpha.Validate(args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.Join(utils.Permutations(args[0]), "\n"))
	},
}

func init() {
	RootCmd.AddCommand(permsCmd)
	permsCmd.Flags().SortFlags = true
}
