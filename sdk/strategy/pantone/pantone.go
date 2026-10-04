package pantone

import (
	"github.com/null93/superperm/sdk/strategy/egan"
	"github.com/null93/superperm/sdk/strategy/standard"
)

var eight = "200020002002000202220002002000202220020002220200020020002220"

var nineA = "20020002000020002000200020002000200200020000200020002000200200002000200020002000" +
	"20020002000020002000200020020000200020002000200020020002000020002000200020002000" +
	"20020002000020002000200020002000200200020000200020002000200020002002000200002000" +
	"20002000200020002002000200002000200020002000200020020002000020002000200020002000" +
	"20020002000200002002000200020000200020002000200020002002000020002000200020002002" +
	"00020000200020002000200020002002000200002000200020002000200020020002000020002000" +
	"20002000200020020002000020002000200020002000200200020002000020020000200200002000" +
	"20002002000020002000"

var nineB = "00200020020002000020002000200020002000200200020000200020020000200020020002000020" +
	"002000200020"

var nineFull = []string{
	"0132456", "0136524", "0142365", "0145623", "0152634", "0154326", "0162543", "0163425",
	"0231546", "0236415", "0241635", "0245316", "0251364", "0254613", "0261453", "0263514",
	"0312564", "0316452", "0321465", "0326541", "0345261", "0346125", "0354162", "0356214",
	"0412653", "0415362", "0421356", "0425631", "0435126", "0436251", "0463152", "0465213",
	"0512346", "0514632", "0521643", "0524361", "0534216", "0536142", "0563241", "0564123",
	"0612435", "0613542", "0621534", "0623451", "0643215", "0645132", "0653124", "0654231",
}

var nineStarts = []string{
	"324560178", "560421378", "421605378", "605123478", "123056478", "056324178",
	"243610578", "610345278", "345106278", "106542378", "542061378", "061243578",
	"612430587", "430521678", "521304678", "304126578", "263401578", "401365278",
	"365014278", "014562378", "562140378", "140263578", "402631587", "631520478",
	"520316478", "316024578", "024163578", "163425078", "425631078", "563124087",
	"312406578", "406215378", "215064378", "064513278", "513640278", "640312578",
	"126043578", "043625178", "625430178", "543026187", "302614578", "614205378",
	"205146378", "146503278", "503461278", "461302578", "023154678", "154326078",
	"326541078", "541620378", "620415378", "415023678",
}

type slice struct {
	x       []uint8
	s       uint8
	classes int
}

type piece struct {
	component int
	start     int
	overlap   int
}

func Generate(alphabet string) string {
	n := len(alphabet)
	if n < 8 {
		result := standard.Generate(alphabet)
		if other := egan.Generate(alphabet); len(other) < len(result) {
			result = other
		}
		return result
	}
	selection, starts := eightSeed(), []string(nil)
	if n == 9 {
		selection, starts = nineSeed(), nineStarts
	}
	for k := len(selection[0].x) + 1; k < n-1; k++ {
		selection = transport(selection, k, uint8(k))
	}
	word := join(complete(selection, n, uint8(n-1)), n, starts)
	mapping := make([]byte, n)
	for i := 0; i < n; i++ {
		mapping[word[i]] = alphabet[i]
	}
	result := make([]byte, len(word))
	for i, c := range word {
		result[i] = mapping[c]
	}
	return string(result)
}

func digits(text string) []uint8 {
	x := make([]uint8, len(text))
	for i, c := range text {
		x[i] = uint8(c - '0')
	}
	return x
}

func walk(start, deficits string) []slice {
	x := digits(start)
	m := len(x)
	selection := []slice{}
	for _, c := range deficits {
		d := int(c - '0')
		selection = append(selection, slice{x, uint8(m), m - d})
		y := make([]uint8, 0, m+2)
		if d == 2 {
			y = append(append(append(y, x[m-1]), x[:m-3]...), x[m-2], x[m-3])
		} else {
			y = append(append(y, x[1:m-1]...), x[0], x[m-1])
		}
		x = y
	}
	return selection
}

func eightSeed() []slice {
	return walk("012345", eight+eight)
}

func nineSeed() []slice {
	selection := append(walk("0123456", nineA), walk("0126435", nineB)...)
	for _, x := range nineFull {
		selection = append(selection, slice{digits(x), 7, 0})
	}
	return selection
}

func insert(x []uint8, i int, letter uint8) []uint8 {
	y := make([]uint8, 0, len(x)+2)
	y = append(y, x[:i]...)
	y = append(y, letter)
	return append(y, x[i:]...)
}

func rotate(x []uint8, i int) []uint8 {
	y := make([]uint8, 0, len(x)+2)
	y = append(y, x[i:]...)
	return append(y, x[:i]...)
}

func transport(selection []slice, k int, w uint8) []slice {
	out := make([]slice, 0, len(selection)*(k-1))
	for _, sl := range selection {
		for i := 0; i < k-1; i++ {
			y := insert(sl.x, i, w)
			if sl.classes == k-3 && i == k-2 {
				y = rotate(y, k-1)
			}
			out = append(out, slice{y, sl.s, sl.classes + 1})
		}
	}
	return out
}

func complete(selection []slice, n int, z uint8) []slice {
	out := []slice{}
	for _, sl := range selection {
		for i := 0; i < n-2 && sl.classes > 0; i++ {
			classes := sl.classes
			if i < sl.classes {
				classes++
			}
			out = append(out, slice{insert(sl.x, i, z), sl.s, classes})
		}
		for i := sl.classes; i <= n-3; i++ {
			out = append(out, slice{append(rotate(sl.x, i), sl.s), z, n - 1})
		}
	}
	return out
}

func (sl slice) word() []uint8 {
	k := len(sl.x) + 1
	p := append(append(make([]uint8, 0, k), sl.x...), sl.s)
	out := append(make([]uint8, 0, (k+1)*sl.classes+k-2), p...)
	for j := 0; j < sl.classes; j++ {
		if j > 0 {
			p = append(append(append(make([]uint8, 0, k), p[2:]...), p[1]), p[0])
			out = append(out, p[k-2], p[k-1])
		}
		for r := 0; r < k-1; r++ {
			p = append(p[1:], p[0])
			out = append(out, p[k-1])
		}
	}
	return out
}

func (sl slice) head(h int) []uint8 {
	return sl.x[:h]
}

func (sl slice) tail(h int) []uint8 {
	y := rotate(sl.x, sl.classes-1)
	return y[len(y)-h:]
}

func encode(word []uint8, n int) uint64 {
	code := uint64(0)
	for _, c := range word {
		code = code*uint64(n) + uint64(c)
	}
	return code
}

func join(slices []slice, n int, starts []string) []uint8 {
	h := n - 3
	ids := map[uint64]int32{}
	words := [][]uint8{}
	vertex := func(word []uint8) int32 {
		code := encode(word, n)
		if id, ok := ids[code]; ok {
			return id
		}
		ids[code] = int32(len(words))
		words = append(words, word)
		return int32(len(words) - 1)
	}
	from := make([]int32, len(slices))
	to := make([]int32, len(slices))
	for e, sl := range slices {
		from[e] = vertex(sl.head(h))
		to[e] = vertex(sl.tail(h))
	}
	first := make([]int32, len(words))
	for v := range first {
		first[v] = -1
	}
	next := make([]int32, len(slices))
	balance := make([]int, len(words))
	for e := len(slices) - 1; e >= 0; e-- {
		next[e] = first[from[e]]
		first[from[e]] = int32(e)
		balance[from[e]]++
		balance[to[e]]--
	}
	for _, b := range balance {
		if b != 0 {
			panic("completed slices are not balanced")
		}
	}
	circuits := [][]int32{}
	for v := range words {
		if first[v] < 0 {
			continue
		}
		stack := []int32{}
		circuit := []int32{}
		at := int32(v)
		for {
			if e := first[at]; e >= 0 {
				first[at] = next[e]
				stack = append(stack, e)
				at = to[e]
			} else {
				if len(stack) == 0 {
					break
				}
				e := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				circuit = append(circuit, e)
				at = from[e]
			}
		}
		for a, b := 0, len(circuit)-1; a < b; a, b = a+1, b-1 {
			circuit[a], circuit[b] = circuit[b], circuit[a]
		}
		circuits = append(circuits, circuit)
	}
	if starts != nil {
		return spell(slices, circuits, starts, n)
	}
	out := []uint8{}
	for _, pc := range greedy(circuits, from, words, n) {
		out = append(out, trail(slices, circuits[pc.component], pc.start, h)[pc.overlap:]...)
	}
	return out
}

func trail(slices []slice, circuit []int32, start, h int) []uint8 {
	out := []uint8{}
	for j := range circuit {
		w := slices[circuit[(start+j)%len(circuit)]].word()
		if j == 0 {
			out = append(out, w...)
		} else {
			out = append(out, w[h:]...)
		}
	}
	return out
}

func isPermutation(window []uint8) bool {
	seen := [256]bool{}
	for _, c := range window {
		if seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}

func open(slices []slice, circuit []int32, position, offset, n int) []uint8 {
	h := n - 3
	z := trail(slices, circuit, position, h)[h:]
	at := func(i int) uint8 {
		return z[((i%len(z))+len(z))%len(z)]
	}
	start := offset - h
	window := make([]uint8, n)
	back := 1
	for ; back <= n; back++ {
		for k := range window {
			window[k] = at(start - back + k)
		}
		if isPermutation(window) {
			break
		}
	}
	out := make([]uint8, len(z)-back+n)
	for i := range out {
		out[i] = at(start + i)
	}
	return out
}

func spell(slices []slice, circuits [][]int32, starts []string, n int) []uint8 {
	component := make([]int, len(slices))
	position := make([]int, len(slices))
	for c, circuit := range circuits {
		for p, e := range circuit {
			component[e], position[e] = c, p
		}
	}
	type location struct{ slice, offset int }
	where := map[string]location{}
	for i, sl := range slices {
		w := sl.word()
		for o := 0; o+n <= len(w); o++ {
			if isPermutation(w[o : o+n]) {
				where[string(w[o:o+n])] = location{i, o}
			}
		}
	}
	used := make([]bool, len(circuits))
	out := []uint8{}
	for _, start := range starts {
		at, ok := where[string(digits(start))]
		if !ok || used[component[at.slice]] {
			panic("start recipe does not match the completed trails")
		}
		used[component[at.slice]] = true
		piece := open(slices, circuits[component[at.slice]], position[at.slice], at.offset, n)
		overlap := n - 1
		for ; overlap > 0; overlap-- {
			if overlap <= len(out) && string(out[len(out)-overlap:]) == string(piece[:overlap]) {
				break
			}
		}
		out = append(out, piece[overlap:]...)
	}
	return out
}

func greedy(circuits [][]int32, from []int32, words [][]uint8, n int) []piece {
	h := n - 3
	index := make([]map[uint64][][2]int32, h)
	for l := 1; l < h; l++ {
		index[l] = map[uint64][][2]int32{}
	}
	for c, circuit := range circuits {
		seen := map[int32]bool{}
		for p, e := range circuit {
			if v := from[e]; !seen[v] {
				seen[v] = true
				for l := 1; l < h; l++ {
					key := encode(words[v][:l], n)
					index[l][key] = append(index[l][key], [2]int32{int32(c), int32(p)})
				}
			}
		}
	}
	used := make([]bool, len(circuits))
	used[0] = true
	order := []piece{{0, 0, 0}}
	end := words[from[circuits[0][0]]]
	fallback := 0
	for len(order) < len(circuits) {
		chosen := piece{-1, 0, 0}
		for l := h - 1; l > 0 && chosen.component < 0; l-- {
			key := encode(end[h-l:], n)
			for len(index[l][key]) > 0 && chosen.component < 0 {
				candidate := index[l][key][0]
				index[l][key] = index[l][key][1:]
				if !used[candidate[0]] {
					chosen = piece{int(candidate[0]), int(candidate[1]), l}
				}
			}
		}
		if chosen.component < 0 {
			for used[fallback] {
				fallback++
			}
			chosen = piece{fallback, 0, 0}
		}
		used[chosen.component] = true
		order = append(order, chosen)
		end = words[from[circuits[chosen.component][chosen.start]]]
	}
	return order
}
