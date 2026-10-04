package utils

import (
	"fmt"
)

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
	if len(perms) == 0 {
		return heatmap
	}
	set := map[string]bool{}
	for _, p := range perms {
		set[p] = true
	}
	n := len(perms[0])
	for i := 0; i+n <= len(input); i++ {
		if set[input[i:i+n]] {
			for j := 0; j < n; j++ {
				heatmap[i+j]++
			}
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
