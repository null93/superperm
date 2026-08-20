package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/null93/superperm/sdk/utils"
	"github.com/spf13/cobra"
)

var analyzeInteractCmd = &cobra.Command{
	Use:     "interact -a ALPHABET -f FILE",
	Short:   "Rotate cycles of a solution in place and watch the layout change",
	Example: "interact -a ABCD -f ./solutions/shortest/4-33.txt",
	Aliases: []string{"inter"},
	Args:    cobra.NoArgs,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		alphabet, _ := cmd.Flags().GetString("alphabet")
		solutionPath, _ := cmd.Flags().GetString("file")
		if len(alphabet) < 1 {
			return fmt.Errorf("alphabet must be at least 1 character long")
		}
		if _, err := os.Stat(solutionPath); err != nil {
			return fmt.Errorf("file does not exist")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		alphabet, _ := cmd.Flags().GetString("alphabet")
		solutionPath, _ := cmd.Flags().GetString("file")
		contents, _ := os.ReadFile(solutionPath)
		permutations := utils.Permutations(alphabet)
		solution := strings.TrimSpace(string(contents))
		if err := interact(permutations, solution); err != nil {
			fmt.Println(err)
		}
	},
}

func interact(perms []string, solution string) error {
	restore, err := rawMode()
	if err != nil {
		return err
	}
	defer restore()
	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")
	cycles := utils.ExtractCycles(perms, solution)
	if len(cycles) < 1 {
		return fmt.Errorf("no cycles found, check the alphabet")
	}
	selectedCycle := 0
	buffer := make([]byte, 32)
	pending := []byte{}
	for {
		if len(pending) < 1 {
			draw(perms, cycles, selectedCycle)
			read, err := os.Stdin.Read(buffer)
			if err != nil || read < 1 {
				return nil
			}
			pending = append(pending, buffer[:read]...)
		}
		key := nextKey(pending)
		pending = pending[len(key):]
		arrow := parseArrow(key)
		switch {
		case key[0] == 'q' || key[0] == 3:
			return nil
		case arrow == 'D':
			if selectedCycle > 0 {
				selectedCycle--
			}
		case arrow == 'C':
			if selectedCycle < len(cycles)-1 {
				selectedCycle++
			}
		case arrow == 'A':
			if selectedCycle > 0 {
				cycles = utils.MoveCycle(cycles, selectedCycle, -1)
				selectedCycle--
			}
		case arrow == 'B':
			if selectedCycle < len(cycles)-1 {
				cycles = utils.MoveCycle(cycles, selectedCycle, 1)
				selectedCycle++
			}
		case key[0] == 'r' || key[0] == 'R':
			cycles[selectedCycle] = utils.RotateElements(cycles[selectedCycle], rotation(key[0] == 'R'))
		case key[0] == 's' || key[0] == 'S':
			direction := rotation(key[0] == 'S')
			cycles = utils.RotateCycles(cycles, direction)
			selectedCycle = utils.RotatedIndex(selectedCycle, len(cycles), direction)
		}
	}
}

func nextKey(pending []byte) []byte {
	if len(pending) > 2 && pending[0] == 27 && pending[1] == '[' {
		for i := 2; i < len(pending); i++ {
			if pending[i] > 0x3f && pending[i] < 0x7f {
				return pending[0 : i+1]
			}
		}
	}
	return pending[0:1]
}

func parseArrow(key []byte) byte {
	if len(key) < 3 || key[0] != 27 || key[1] != '[' {
		return 0
	}
	final := key[len(key)-1]
	switch final {
	case 'a', 'b', 'c', 'd':
		return final - 32
	case 'A', 'B', 'C', 'D':
		return final
	}
	return 0
}

func rotation(clockwise bool) int {
	if clockwise {
		return 1
	}
	return -1
}

func draw(perms []string, cycles []utils.Cycle, selected int) {
	solution, layout := utils.RenderCycles(cycles, selected)
	status := "invalid"
	if utils.Validate(perms, solution) {
		status = "valid"
	}
	lines := []string{
		fmt.Sprintf("%d", len(solution)),
		solution,
	}
	lines = append(lines, strings.Split(strings.TrimRight(layout, "\n"), "\n")...)
	lines = append(lines, "", fmt.Sprintf("cycle %d/%d %s | left/right select | up/down move | r/R rotate elements | s/S rotate cycles | q quit",
		selected+1, len(cycles), status))
	fmt.Print("\033[H\033[2J")
	for _, line := range lines {
		fmt.Print(line + "\r\n")
	}
}

func rawMode() (func(), error) {
	enable := exec.Command("stty", "raw", "-echo")
	enable.Stdin = os.Stdin
	if err := enable.Run(); err != nil {
		return nil, fmt.Errorf("cannot read keys, stdin is not a terminal")
	}
	return func() {
		disable := exec.Command("stty", "sane")
		disable.Stdin = os.Stdin
		disable.Run()
	}, nil
}

func init() {
	analyzeCmd.AddCommand(analyzeInteractCmd)
	analyzeInteractCmd.Flags().SortFlags = true
	analyzeInteractCmd.Flags().StringP("alphabet", "a", "", "exact alphabet used to generate the solution")
	analyzeInteractCmd.Flags().StringP("file", "f", "", "path to file with solution")
}
