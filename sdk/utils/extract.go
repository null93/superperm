package utils

import (
	"fmt"
	"html"
	"os"
	"strings"
)

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

func RenderSVG(rows [][]Extraction) string {
	fontSize := 15
	charWidth := 9
	lineHeight := 20
	padding := 16
	columns := 0
	for _, row := range rows {
		for _, extraction := range row {
			if end := extraction.Index + len(extraction.Permutation); end > columns {
				columns = end
			}
		}
	}
	width := columns*charWidth + padding*2
	height := len(rows)*lineHeight + padding*2
	output := strings.Builder{}
	fmt.Fprintf(&output, "<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"%d\" height=\"%d\" viewBox=\"0 0 %d %d\">\n", width, height, width, height)
	fmt.Fprintf(&output, "<style>text{font-family:ui-monospace,SFMono-Regular,Menlo,Consolas,\"DejaVu Sans Mono\",\"Liberation Mono\",monospace;font-size:%dpx;fill:%s;white-space:pre}", fontSize, HexForeground)
	for i, hex := range HexColors {
		fmt.Fprintf(&output, ".c%d{fill:%s}", i, hex)
	}
	output.WriteString("</style>\n")
	fmt.Fprintf(&output, "<rect width=\"100%%\" height=\"100%%\" fill=\"%s\"/>\n", HexBackground)
	for r, row := range rows {
		y := padding + r*lineHeight + (lineHeight+fontSize)/2 - 2
		for _, extraction := range row {
			x := padding + extraction.Index*charWidth
			length := len(extraction.Permutation) * charWidth
			class := ""
			if colorEnabled {
				class = fmt.Sprintf(" class=\"c%d\"", ColorIndex(extraction.Color))
			}
			fmt.Fprintf(&output, "<text%s x=\"%d\" y=\"%d\" textLength=\"%d\" lengthAdjust=\"spacing\">%s</text>\n", class, x, y, length, html.EscapeString(extraction.Permutation))
		}
	}
	output.WriteString("</svg>\n")
	return output.String()
}

func PrintExtraction(perms []string, superpermutation string) {
	fmt.Print(RenderRows(PackExtractions(Extract(perms, superpermutation)), -1))
}

func WriteExtraction(perms []string, superpermutation, path string) error {
	return os.WriteFile(path, []byte(RenderSVG(PackExtractions(Extract(perms, superpermutation)))), 0644)
}
