package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/null93/superperm/sdk/alpha"
	"github.com/null93/superperm/sdk/utils"
	"github.com/spf13/cobra"
)

var utilityPermutationsCmd = &cobra.Command{
	Use:     "permutations -a ALPHABET -n INT",
	Short:   "generate all permutations of an alphabet",
	Example: "permutations -a latin -n 4",
	Aliases: []string{"perm", "perms"},
	Args:    cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		n, _ := cmd.Flags().GetInt("size")
		output, _ := cmd.Flags().GetString("output")
		alphabetName, _ := cmd.Flags().GetString("alphabet")
		if n < 1 {
			return alpha.ErrAlphabetSize
		}
		if _, errAlpha := alpha.GetAlphabet(alphabetName, n); errAlpha != nil {
			return errAlpha
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
		n, _ := cmd.Flags().GetInt("size")
		output, _ := cmd.Flags().GetString("output")
		alphabetName, _ := cmd.Flags().GetString("alphabet")
		alphabet, _ := alpha.GetAlphabet(alphabetName, n)
		permutations := utils.Permutations(alphabet)
		if output != "" {
			os.MkdirAll(filepath.Dir(output), 0755)
			file, _ := os.Create(output)
			defer file.Close()
			file.WriteString(strings.Join(permutations, "\n"))
		} else {
			fmt.Println(strings.Join(permutations, "\n"))
		}
	},
}

func init() {
	utilityCmd.AddCommand(utilityPermutationsCmd)
	utilityPermutationsCmd.Flags().SortFlags = true
	utilityPermutationsCmd.Flags().IntP("size", "n", 0, "alphabet size")
	utilityPermutationsCmd.Flags().StringP("alphabet", "a", "latin", "which alphabet to use")
	utilityPermutationsCmd.Flags().StringP("output", "o", "", "path to solutions directory")
}
