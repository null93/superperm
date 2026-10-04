package utils

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/null93/superperm/sdk/alpha"
)

func ReadSolution(path string) (string, string, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return "", "", fmt.Errorf("file does not exist")
	}
	solution := strings.TrimSpace(string(contents))
	alphabet := ExtractAlphabet(solution)
	if len(alphabet) < 1 {
		return "", "", fmt.Errorf("file is empty")
	}
	if errAlpha := alpha.Validate(alphabet); errAlpha != nil {
		return "", "", fmt.Errorf("solution %v", errAlpha)
	}
	return solution, alphabet, nil
}

func ExtractAlphabet(input string) string {
	seen := map[rune]bool{}
	alphabet := []rune{}
	for _, r := range input {
		if !seen[r] {
			seen[r] = true
			alphabet = append(alphabet, r)
		}
	}
	sort.Slice(alphabet, func(i, j int) bool {
		return alphabet[i] < alphabet[j]
	})
	return string(alphabet)
}

func windows(input string, n int) map[string]bool {
	found := map[string]bool{}
	for i := 0; i+n <= len(input); i++ {
		found[input[i:i+n]] = true
	}
	return found
}

func MissingPermutations(alphabet, input string) []string {
	found := windows(input, len(alphabet))
	missing := []string{}
	for _, p := range Permutations(alphabet) {
		if !found[p] {
			missing = append(missing, p)
		}
	}
	return missing
}

func IsSuperpermutation(perms []string, input string) bool {
	if len(perms) == 0 {
		return true
	}
	found := windows(input, len(perms[0]))
	for _, p := range perms {
		if !found[p] {
			return false
		}
	}
	return true
}
