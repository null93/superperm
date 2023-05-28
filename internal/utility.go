package internal

import (
	"github.com/spf13/cobra"
)

var utilityCmd = &cobra.Command{
	Use:     "utility",
	Short:   "Helpful utilities",
	Aliases: []string{"util", "utils"},
}

func init() {
	RootCmd.AddCommand(utilityCmd)
	utilityCmd.Flags().SortFlags = true
}
