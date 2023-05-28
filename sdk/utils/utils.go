package utils

import (
	"fmt"
	"strings"
)

var HistogramPoint = "\033[33m♦\033[0m"

func DisableColor(state bool) {
	if state {
		HistogramPoint = "♦"
	} else {
		HistogramPoint = "\033[33m♦\033[0m"
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

func Validate(set []string, input string) bool {
	for _, p := range set {
		if strings.Index(input, p) == -1 {
			return false
		}
	}
	return true
}

func PrintHistogram(alphabet, input string) {
	n := len(alphabet)
	mapping := map[rune]int{}
	for i, r := range alphabet {
		mapping[r] = i + 1
	}
	indexed := []int{}
	for _, r := range input {
		fmt.Printf("%c ", r)
		indexed = append(indexed, mapping[r])
	}
	fmt.Println()
	for i := 0; i < n; i++ {
		for j, r := range indexed {
			if r > 0 {
				fmt.Print(HistogramPoint + " ")
				indexed[j]--
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Println()
	}
}

func GetHeatMap(perms []string, input string) []uint8 {
	heatmap := make([]uint8, len(input))
	for _, p := range perms {
		target := strings.Clone(input)
		n := len(p)
		i := strings.Index(target, p)
		seen := 0
		for i > -1 {
			for j := 0; j < n; j++ {
				heatmap[seen+i+j]++
			}
			seen += i + n
			target = target[i+len(p):]
			i = strings.Index(target, p)
		}
	}
	return heatmap
}

func PrintHeatMap(perms []string, input string) {
	heatmap := GetHeatMap(perms, input)
	max := uint8(0)
	for j, _ := range input {
		fmt.Printf("%c ", input[j])
	}
	fmt.Println()
	for j, _ := range heatmap {
		fmt.Printf("%d ", heatmap[j])
		if heatmap[j] > max {
			max = heatmap[j]
		}
	}
	fmt.Println()
	for i := uint8(0); i < max; i++ {
		for j, _ := range heatmap {
			if heatmap[j] > 0 {
				fmt.Print(HistogramPoint + " ")
				heatmap[j]--
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Println()
	}
}

func Extract(perms []string, superpermutation string) []string {
	results := []string{}
	for _, p := range perms {
		n := len(p)
		last := 0
		layer := strings.Repeat(" ", len(superpermutation))
		index := strings.Index(superpermutation[last:], p)
		for index != -1 {
			layer = layer[0:index+last] + p + layer[index+last+n:]
			last = index + n
			fmt.Println(last, superpermutation[last:])
			index = strings.Index(superpermutation[last:], p)
			fmt.Println("index", index)
		}
		for _, r := range layer {
			fmt.Printf("%c ", r)
		}
		fmt.Println()
		results = append(results, layer)
	}
	return results
}

func PrintExtraction() {

}

func Factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}
