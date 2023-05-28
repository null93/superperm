package internal

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "superperm",
	Version: "2.0.0",
}

func init() {
	RootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	RootCmd.Flags().SortFlags = true
}
