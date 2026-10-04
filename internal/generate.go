package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/null93/superperm/sdk/alpha"
	"github.com/null93/superperm/sdk/strategy"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate ALPHABET [flags]",
	Aliases: []string{"gen"},
	Short:   "Generate a superpermutation",
	Args:    cobra.ExactArgs(1),
	Example: "generate ABCD\nsuperperm perms ABCD | xargs -n1 superperm generate",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		strategyName, _ := cmd.Flags().GetString("strategy")
		output, _ := cmd.Flags().GetString("output")
		if errAlpha := alpha.Validate(args[0]); errAlpha != nil {
			return errAlpha
		}
		if _, errStrategy := strategy.Get(strategyName); errStrategy != nil {
			return fmt.Errorf("%v, use one of: %s", errStrategy, strings.Join(strategy.Names(), ", "))
		}
		if output != "" {
			info, err := os.Stat(output)
			if os.IsNotExist(err) {
				return fmt.Errorf("directory does not exist")
			}
			if !info.IsDir() {
				return fmt.Errorf("not a directory")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		strategyName, _ := cmd.Flags().GetString("strategy")
		output, _ := cmd.Flags().GetString("output")
		alphabet := args[0]
		generate, _ := strategy.Get(strategyName)
		solution := generate(alphabet)
		if output != "" {
			path := filepath.Join(output, strategyName, fmt.Sprintf("%d-%d.txt", len(alphabet), len(solution)))
			os.MkdirAll(filepath.Dir(path), 0755)
			if err := os.WriteFile(path, []byte(solution), 0644); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
		fmt.Println(solution)
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)
	generateCmd.Flags().SortFlags = true
	generateCmd.Flags().StringP("strategy", "s", "standard", "which strategy to use, one of: "+strings.Join(strategy.Names(), ", "))
	generateCmd.Flags().StringP("output", "o", "", "path to solutions directory")
}
