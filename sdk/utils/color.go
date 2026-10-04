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
