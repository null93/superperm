package standard

import (
	"github.com/null93/superperm/sdk/utils"
)

func Generate(alphabet string) string {
	return generate(alphabet, 1)
}

func rotations(seed string, level int) []string {
	result := []string{}
	for i := 0; i < level; i++ {
		rotation := utils.RotateString(seed, 0, level-1, i)
		result = append(result, rotation)
	}
	return result
}

func generate(seed string, level int) string {
	n := len(seed)
	if level > n || level < 1 {
		return ""
	}
	if level == n {
		result := ""
		for r, rotation := range rotations(seed, level) {
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
	for r, rotation := range rotations(seed, level) {
		collected := generate(rotation, level+1)
		if r == level-1 {
			results += collected
		} else {
			m := len(collected)
			results += collected[0 : m-level+1]
		}
	}
	return results
}
