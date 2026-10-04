package utils

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
var HexBackground = "#1e1e1e"
var HexForeground = "#cccccc"
var HexColors = []string{
	"#cd3131",
	"#0dbc79",
	"#e5e510",
	"#2472c8",
	"#bc3fbc",
	"#11a8cd",
	"#f14c4c",
	"#23d18b",
	"#f5f543",
	"#3b8eea",
	"#d670d6",
	"#29b8db",
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
	return Colors[ColorIndex(i)] + input + ColorReset
}

func ColorIndex(i int) int {
	n := len(Colors)
	return ((i % n) + n) % n
}
