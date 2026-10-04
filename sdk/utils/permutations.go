package utils

import (
	"fmt"
	"strings"
)

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Permutations(alphabet string) []string {
	n := len(alphabet)
	if n == 0 {
		return []string{}
	} else if n == 1 {
		return []string{alphabet}
	} else {
		result := []string{}
		for i, r := range alphabet {
			rest := alphabet[0:i] + alphabet[i+1:]
			for _, p := range Permutations(rest) {
				result = append(result, string(r)+p)
			}
		}
		return result
	}
}

func RotateArray(input []string, start, end, rotations int) []string {
	if start < 0 || end >= len(input) || start > end {
		panic(fmt.Sprintf("invalid range: start %d; ends %d; length %d", start, end, len(input)))
	}
	if rotations == 0 || end-start < 1 {
		return input
	}
	left := append([]string{}, input[0:start]...)
	center := append([]string{}, input[start:end+1]...)
	right := append([]string{}, input[end+1:]...)
	n := end - start + 1
	for r := 0; r < rotations%n; r++ {
		center = append(center[1:], center[0])
	}
	return append(left, append(center, right...)...)
}

func RotateString(input string, start, end, rotations int) string {
	array := strings.Split(input, "")
	result := RotateArray(array, start, end, rotations)
	return strings.Join(result, "")
}
