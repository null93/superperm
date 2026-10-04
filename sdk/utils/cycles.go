package utils

import "strings"

type Cycle struct {
	Color    int
	Elements []string
}

func ExtractCycles(perms []string, superpermutation string) []Cycle {
	cycles := []Cycle{}
	for c, thread := range GroupThreads(Extract(perms, superpermutation)) {
		elements := []string{}
		for _, extraction := range thread {
			elements = append(elements, extraction.Permutation)
		}
		cycles = append(cycles, Cycle{c, elements})
	}
	return cycles
}

func MergeCycles(cycles []Cycle) (string, [][]int) {
	merged := ""
	positions := [][]int{}
	for _, cycle := range cycles {
		indexes := []int{}
		for _, element := range cycle.Elements {
			if merged == "" {
				indexes = append(indexes, 0)
				merged = element
				continue
			}
			overlap := len(element) - 1
			if len(merged) < overlap {
				overlap = len(merged)
			}
			for ; overlap > 0; overlap-- {
				if strings.HasSuffix(merged, element[0:overlap]) {
					break
				}
			}
			indexes = append(indexes, len(merged)-overlap)
			merged += element[overlap:]
		}
		positions = append(positions, indexes)
	}
	return merged, positions
}

func CycleExtractions(cycles []Cycle, positions [][]int) []Extraction {
	extractions := []Extraction{}
	for c, cycle := range cycles {
		for e, element := range cycle.Elements {
			extractions = append(extractions, Extraction{positions[c][e], element, c, cycle.Color})
		}
	}
	return extractions
}

func MoveCycle(cycles []Cycle, index, direction int) []Cycle {
	target := index + direction
	if index < 0 || index > len(cycles)-1 || target < 0 || target > len(cycles)-1 {
		return cycles
	}
	moved := append([]Cycle{}, cycles...)
	moved[index], moved[target] = moved[target], moved[index]
	return moved
}

func RotateCycles(cycles []Cycle, direction int) []Cycle {
	if len(cycles) < 2 {
		return cycles
	}
	rotated := append([]Cycle{}, cycles...)
	if direction < 0 {
		return append(rotated[len(rotated)-1:], rotated[0:len(rotated)-1]...)
	}
	return append(rotated[1:], rotated[0])
}

func RotatedIndex(index, count, direction int) int {
	if count < 2 {
		return index
	}
	if direction < 0 {
		return (index + 1) % count
	}
	return ((index-1)%count + count) % count
}

func RotateElements(cycle Cycle, direction int) Cycle {
	if len(cycle.Elements) < 2 {
		return cycle
	}
	rotations := 1
	if direction < 0 {
		rotations = len(cycle.Elements) - 1
	}
	elements := append([]string{}, cycle.Elements...)
	return Cycle{cycle.Color, RotateArray(elements, 0, len(elements)-1, rotations)}
}

func RenderCycles(cycles []Cycle, selected int) (string, string) {
	merged, positions := MergeCycles(cycles)
	rows := PackExtractions(CycleExtractions(cycles, positions))
	return merged, RenderRows(rows, selected)
}
