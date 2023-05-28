package internal

import (
	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze superpermutation solutions",
}

func init() {
	RootCmd.AddCommand(analyzeCmd)
	analyzeCmd.Flags().SortFlags = true
}
