package rotate

import (
	"time"

	"github.com/null93/superperm/sdk/utils"
)

type Problem struct {
	Alphabet string
	Solution string
	Length   int
	Duration time.Duration
}

func GetRotations(seed string, level int) []string {
	rotations := []string{}
	for i := 0; i < level; i++ {
		rotation := utils.RotateString(seed, 0, level-1, i)
		rotations = append(rotations, rotation)
	}
	return rotations
}

func GeneratePermutations(seed string, level int) []string {
	n := len(seed)
	if level > n || level < 1 {
		return []string{}
	}
	if level == n {
		return GetRotations(seed, level)
	}
	results := []string{}
	for _, rotation := range GetRotations(seed, level) {
		collected := GeneratePermutations(rotation, level+1)
		results = append(results, collected...)
	}
	return results
}

func GenerateSuperPermutation(seed string, level int) string {
	n := len(seed)
	if level > n || level < 1 {
		return ""
	}
	if level == n {
		result := ""
		for r, rotation := range GetRotations(seed, level) {
			if r == level-1 {
				result += rotation
			} else {
				m := len(rotation)
				result += rotation[0 : m-level+1]
			}
		}
		return result
	}
	results := ""
	for r, rotation := range GetRotations(seed, level) {
		collected := GenerateSuperPermutation(rotation, level+1)
		if r == level-1 {
			results += collected
		} else {
			m := len(collected)
			results += collected[0 : m-level+1]
		}
	}
	return results
}

func (p *Problem) Solve(alphabet string) {
	start := time.Now()
	p.Solution = GenerateSuperPermutation(alphabet, 1)
	p.Duration = time.Since(start)
	p.Length = len(p.Solution)
	p.Alphabet = alphabet
}

func (p *Problem) GetSolution() string {
	return p.Solution
}

func (p *Problem) GetLength() int {
	return len(p.Solution)
}

func (p *Problem) GetDuration() time.Duration {
	return p.Duration
}
