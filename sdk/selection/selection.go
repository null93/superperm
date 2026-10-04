package selection

type Slice struct {
	X       []uint8
	S       uint8
	Classes int
}

type piece struct {
	component int
	start     int
	overlap   int
}

func Digits(text string) []uint8 {
	x := make([]uint8, len(text))
	for i, c := range text {
		x[i] = uint8(c - '0')
	}
	return x
}

func Walk(start, deficits string) []Slice {
	x := Digits(start)
	m := len(x)
	selection := []Slice{}
	for _, c := range deficits {
		d := int(c - '0')
		selection = append(selection, Slice{x, uint8(m), m - d})
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

func Transport(selection []Slice, k int, w uint8) []Slice {
	out := make([]Slice, 0, len(selection)*(k-1))
	for _, sl := range selection {
		for i := 0; i < k-1; i++ {
			y := insert(sl.X, i, w)
			if sl.Classes == k-3 && i == k-2 {
				y = rotate(y, k-1)
			}
			out = append(out, Slice{y, sl.S, sl.Classes + 1})
		}
	}
	return out
}

func Complete(selection []Slice, n int, z uint8) []Slice {
	out := []Slice{}
	for _, sl := range selection {
		for i := 0; i < n-2 && sl.Classes > 0; i++ {
			classes := sl.Classes
			if i < sl.Classes {
				classes++
			}
			out = append(out, Slice{insert(sl.X, i, z), sl.S, classes})
		}
		for i := sl.Classes; i <= n-3; i++ {
			out = append(out, Slice{append(rotate(sl.X, i), sl.S), z, n - 1})
		}
	}
	return out
}

func (sl Slice) word() []uint8 {
	k := len(sl.X) + 1
	p := append(append(make([]uint8, 0, k), sl.X...), sl.S)
	out := append(make([]uint8, 0, (k+1)*sl.Classes+k-2), p...)
	for j := 0; j < sl.Classes; j++ {
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

func (sl Slice) head(h int) []uint8 {
	return sl.X[:h]
}

func (sl Slice) tail(h int) []uint8 {
	y := rotate(sl.X, sl.Classes-1)
	return y[len(y)-h:]
}

func encode(word []uint8, n int) uint64 {
	code := uint64(0)
	for _, c := range word {
		code = code*uint64(n) + uint64(c)
	}
	return code
}

func Join(slices []Slice, n int, starts []string) []uint8 {
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

func trail(slices []Slice, circuit []int32, start, h int) []uint8 {
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

func open(slices []Slice, circuit []int32, position, offset, n int) []uint8 {
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

func spell(slices []Slice, circuits [][]int32, starts []string, n int) []uint8 {
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
		at, ok := where[string(Digits(start))]
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

func Relabel(word []uint8, alphabet string) string {
	mapping := make([]byte, len(alphabet))
	for i := 0; i < len(alphabet); i++ {
		mapping[word[i]] = alphabet[i]
	}
	result := make([]byte, len(word))
	for i, c := range word {
		result[i] = mapping[c]
	}
	return string(result)
}
