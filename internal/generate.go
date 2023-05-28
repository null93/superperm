package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/null93/superperm/sdk/alpha"
	"github.com/null93/superperm/sdk/problem"
	"github.com/null93/superperm/sdk/utils"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate N [flags]",
	Aliases: []string{"gen"},
	Short:   "Generate minimal superpermutation",
	Args:    cobra.ExactArgs(1),
	Example: "generate 3 -a digit -m rotate",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		seedIndex, _ := cmd.Flags().GetInt("seed-index")
		strategy, _ := cmd.Flags().GetString("strategy")
		output, _ := cmd.Flags().GetString("output")
		alphabetName, _ := cmd.Flags().GetString("alphabet")
		n, nErr := strconv.Atoi(args[0])
		if nErr != nil || n < 1 {
			return alpha.ErrAlphabetSize
		}
		if _, errAlpha := alpha.GetAlphabet(alphabetName, n); errAlpha != nil {
			return errAlpha
		}
		nFactorial := utils.Factorial(n)
		if seedIndex < 0 || seedIndex > nFactorial-1 {
			return fmt.Errorf("seed index must be between 0 and %d", nFactorial-1)
		}
		if !problem.IsSupportedStrategy(strategy) {
			return problem.ErrUnsupportedStrategy
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
		n, _ := strconv.Atoi(args[0])
		seedIndex, _ := cmd.Flags().GetInt("seed-index")
		stats, _ := cmd.Flags().GetBool("stats")
		strategy, _ := cmd.Flags().GetString("strategy")
		output, _ := cmd.Flags().GetString("output")
		alphabetName, _ := cmd.Flags().GetString("alphabet")
		alphabet, _ := alpha.GetAlphabet(alphabetName, n)
		permutations := utils.Permutations(alphabet)
		solution, _ := problem.NewProblem(strategy)
		solution.Solve(permutations[seedIndex])
		if output != "" {
			fillPath := filepath.Join(output, strategy, fmt.Sprintf("%d-%d.txt", n, solution.GetLength()))
			os.MkdirAll(filepath.Dir(fillPath), 0755)
			file, _ := os.Create(fillPath)
			defer file.Close()
			file.WriteString(solution.GetSolution())
		}
		if stats {
			fmt.Printf("Length:    %d\n", solution.GetLength())
			fmt.Printf("Duration:  %s\n", solution.GetDuration())
		} else {
			fmt.Println(solution.GetSolution())
		}
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)
	generateCmd.Flags().SortFlags = true
	generateCmd.Flags().Bool("stats", false, "only show stats")
	generateCmd.Flags().StringP("strategy", "s", "shortest", "which strategy to use")
	generateCmd.Flags().StringP("alphabet", "a", "latin", "which alphabet to use")
	generateCmd.Flags().StringP("output", "o", "", "path to solutions directory")
	generateCmd.Flags().IntP("seed-index", "i", 0, "permutation index for initial alphabet")
}
