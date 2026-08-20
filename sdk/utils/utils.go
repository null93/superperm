package utils

import (
	"fmt"
	"strings"
)

var HistogramPoint = "\033[33m♦\033[0m"

var HighlightPoint = "\033[7m"
var ColorReset = "\033[0m"
var Colors = []string{
	"\033[31m",
	"\033[32m",
	"\033[33m",
	"\033[34m",
	"\033[35m",
	"\033[36m",
	"\033[91m",
	"\033[92m",
	"\033[93m",
	"\033[94m",
	"\033[95m",
	"\033[96m",
}

var colorEnabled = true

func DisableColor(state bool) {
	colorEnabled = !state
	if state {
		HistogramPoint = "♦"
	} else {
		HistogramPoint = "\033[33m♦\033[0m"
	}
}

func Colorize(input string, i int) string {
	if !colorEnabled {
		return input
	}
	n := len(Colors)
	return Colors[((i%n)+n)%n] + input + ColorReset
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

type Extraction struct {
	Index       int
	Permutation string
	Thread      int
	Color       int
}

func Extract(perms []string, superpermutation string) []Extraction {
	results := []Extraction{}
	if len(perms) == 0 {
		return results
	}
	set := map[string]bool{}
	for _, p := range perms {
		set[p] = true
	}
	n := len(perms[0])
	thread := -1
	previous := -2
	shift := 0
	for i := 0; i+n <= len(superpermutation); i++ {
		window := superpermutation[i : i+n]
		if set[window] {
			if i != previous+1 {
				thread++
				if thread > 0 {
					shift += i - previous - 1
				}
			}
			results = append(results, Extraction{i, window, thread, shift})
			previous = i
		}
	}
	return results
}

func GroupThreads(extractions []Extraction) [][]Extraction {
	threads := [][]Extraction{}
	for _, extraction := range extractions {
		last := len(threads) - 1
		if last < 0 || threads[last][0].Thread != extraction.Thread {
			threads = append(threads, []Extraction{extraction})
		} else {
			threads[last] = append(threads[last], extraction)
		}
	}
	return threads
}

func PackExtractions(extractions []Extraction) [][]Extraction {
	threads := GroupThreads(extractions)
	if len(threads) < 1 {
		return [][]Extraction{}
	}
	placed := map[int][]Extraction{}
	ends := map[int]int{}
	direction := 1
	last := 0
	lowest := 0
	highest := 0
	for t, thread := range threads {
		if t > 1 {
			delta := thread[0].Index - threads[t-1][len(threads[t-1])-1].Index
			previous := threads[t-1][0].Index - threads[t-2][len(threads[t-2])-1].Index
			if delta > previous {
				direction = 1
			} else if delta < previous {
				direction = -1
			}
		}
		base := 0
		if t > 0 {
			base = last + direction
		}
		for !fitsThread(thread, ends, base, direction) {
			base += direction
		}
		for i, extraction := range thread {
			row := base + direction*i
			placed[row] = append(placed[row], extraction)
			ends[row] = extraction.Index + len(extraction.Permutation)
			if row < lowest {
				lowest = row
			}
			if row > highest {
				highest = row
			}
		}
		last = base + direction*(len(thread)-1)
	}
	rows := [][]Extraction{}
	for row := lowest; row <= highest; row++ {
		if len(placed[row]) > 0 {
			rows = append(rows, placed[row])
		}
	}
	return rows
}

func fitsThread(thread []Extraction, ends map[int]int, base, direction int) bool {
	for i, extraction := range thread {
		if end, taken := ends[base+direction*i]; taken && end > extraction.Index {
			return false
		}
	}
	return true
}

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

func RenderRows(rows [][]Extraction, selected int) string {
	output := strings.Builder{}
	for _, row := range rows {
		column := 0
		for _, extraction := range row {
			output.WriteString(strings.Repeat(" ", extraction.Index-column))
			if extraction.Thread == selected {
				output.WriteString(HighlightPoint)
			}
			output.WriteString(Colorize(extraction.Permutation, extraction.Color))
			column = extraction.Index + len(extraction.Permutation)
		}
		output.WriteString("\n")
	}
	return output.String()
}

func RenderCycles(cycles []Cycle, selected int) (string, string) {
	merged, positions := MergeCycles(cycles)
	rows := PackExtractions(CycleExtractions(cycles, positions))
	return merged, RenderRows(rows, selected)
}

func PrintExtraction(perms []string, superpermutation string) {
	fmt.Println(superpermutation)
	fmt.Print(RenderRows(PackExtractions(Extract(perms, superpermutation)), -1))
}

func Factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}
