package euler

import (
	"github.com/null93/superperm/sdk/selection"
	"github.com/null93/superperm/sdk/strategy/egan"
	"github.com/null93/superperm/sdk/strategy/standard"
)

const budget = 20000

func Generate(alphabet string) string {
	n := len(alphabet)
	if n < 8 {
		return fallback(alphabet)
	}
	deficits := newSpace(6).seed()
	if deficits == "" {
		return fallback(alphabet)
	}
	seed := selection.Walk("012345", deficits)
	for k := 7; k < n-1; k++ {
		seed = selection.Transport(seed, k, uint8(k))
	}
	word := selection.Join(selection.Complete(seed, n, uint8(n-1)), n, nil)
	return selection.Relabel(word, alphabet)
}

func fallback(alphabet string) string {
	result := standard.Generate(alphabet)
	if other := egan.Generate(alphabet); len(other) < len(result) {
		result = other
	}
	return result
}

type space struct {
	m       int
	states  [][]uint8
	index   map[string]int
	full    []int
	short   []int
	class   []int
	members [][]int
	preds   []int
}

func newSpace(m int) *space {
	sp := &space{m: m, index: map[string]int{}}
	var permute func(prefix []uint8, used int)
	permute = func(prefix []uint8, used int) {
		if len(prefix) == m {
			sp.index[string(prefix)] = len(sp.states)
			sp.states = append(sp.states, append([]uint8{}, prefix...))
			return
		}
		for c := 0; c < m; c++ {
			if used&(1<<c) == 0 {
				permute(append(prefix, uint8(c)), used|1<<c)
			}
		}
	}
	permute(nil, 0)
	classes := map[string]int{}
	for i, x := range sp.states {
		f := append(append([]uint8{}, x[1:m-1]...), x[0], x[m-1])
		s := append(append(append([]uint8{x[m-1]}, x[:m-3]...), x[m-2]), x[m-3])
		sp.full = append(sp.full, sp.index[string(f)])
		sp.short = append(sp.short, sp.index[string(s)])
		least := ""
		for r := 0; r < m; r++ {
			if y := string(x[r:]) + string(x[:r]); least == "" || y < least {
				least = y
			}
		}
		if _, ok := classes[least]; !ok {
			classes[least] = len(sp.members)
			sp.members = append(sp.members, nil)
		}
		sp.class = append(sp.class, classes[least])
		sp.members[classes[least]] = append(sp.members[classes[least]], i)
	}
	sp.preds = make([]int, len(sp.members))
	for i := range sp.states {
		sp.preds[sp.class[sp.full[i]]]++
		sp.preds[sp.class[sp.short[i]]]++
	}
	return sp
}

func (sp *space) seed() string {
	best, deficits := len(sp.members)/2+1, ""
	for _, perm := range involutions(sp.m) {
		if s := sp.search(perm, best-1); s != nil && s.found {
			best, deficits = s.shorts, string(s.digits)+string(s.digits)
		}
	}
	return deficits
}

func involutions(m int) [][]uint8 {
	out := [][]uint8{}
	var build func(p []uint8, i int)
	build = func(p []uint8, i int) {
		for i < m && p[i] != 255 {
			i++
		}
		if i == m {
			out = append(out, append([]uint8{}, p...))
			return
		}
		q := append([]uint8{}, p...)
		q[i] = uint8(i)
		build(q, i+1)
		for j := i + 1; j < m; j++ {
			if p[j] == 255 {
				q := append([]uint8{}, p...)
				q[i], q[j] = uint8(j), uint8(i)
				build(q, i+1)
			}
		}
	}
	p := make([]uint8, m)
	for i := range p {
		p[i] = 255
	}
	build(p, 0)
	return out[1:]
}

type search struct {
	sp       *space
	mirror   []int
	maxShort int
	visited  []bool
	avail    []int
	alive    []bool
	endIn    int
	end      int
	path     []int
	digits   []byte
	shorts   int
	run      int
	nodes    int
	found    bool
	exceeded bool
	undo     []int
}

func (sp *space) search(perm []uint8, maxShort int) *search {
	s := &search{
		sp:       sp,
		maxShort: maxShort,
		visited:  make([]bool, len(sp.members)),
		avail:    append([]int{}, sp.preds...),
		alive:    make([]bool, 2*len(sp.states)),
		mirror:   make([]int, len(sp.states)),
	}
	for i, x := range sp.states {
		y := make([]uint8, len(x))
		for j := range x {
			y[j] = perm[x[j]]
		}
		s.mirror[i] = sp.index[string(y)]
	}
	s.end = s.mirror[0]
	if sp.class[0] == sp.class[s.end] {
		return nil
	}
	for p := range sp.states {
		s.alive[2*p], s.alive[2*p+1] = true, true
		if sp.full[p] == s.end {
			s.endIn++
		}
		if sp.short[p] == s.end {
			s.endIn++
		}
	}
	s.path = []int{0}
	s.run = 1
	s.claim(0)
	s.claim(s.end)
	s.undo = s.undo[:0]
	s.dfs()
	return s
}

func (s *search) target(p, d int) int {
	if d == 2 {
		return s.sp.short[p]
	}
	return s.sp.full[p]
}

func (s *search) kill(p, d int) {
	k := 2*p + d/2
	if !s.alive[k] {
		return
	}
	s.alive[k] = false
	t := s.target(p, d)
	s.avail[s.sp.class[t]]--
	if t == s.end {
		s.endIn--
	}
	s.undo = append(s.undo, k)
}

func (s *search) restore(mark int) {
	for len(s.undo) > mark {
		k := s.undo[len(s.undo)-1]
		s.undo = s.undo[:len(s.undo)-1]
		s.alive[k] = true
		t := s.target(k/2, 2*(k%2))
		s.avail[s.sp.class[t]]++
		if t == s.end {
			s.endIn++
		}
	}
}

func (s *search) claim(q int) {
	s.visited[s.sp.class[q]] = true
	for _, r := range s.sp.members[s.sp.class[q]] {
		if r != q {
			s.kill(r, 0)
			s.kill(r, 2)
		}
	}
}

func (s *search) feasible() bool {
	if s.endIn == 0 {
		return false
	}
	for c := range s.avail {
		if !s.visited[c] && s.avail[c] == 0 {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func (s *search) dfs() {
	s.nodes++
	if s.found || s.exceeded {
		return
	}
	if s.nodes > budget {
		s.exceeded = true
		return
	}
	cur := s.path[len(s.path)-1]
	half := len(s.sp.members) / 2
	limit := s.sp.m - 1
	if len(s.path) == half {
		for _, d := range []int{0, 2} {
			t := s.shorts + d/2
			if s.target(cur, d) == s.end && t <= s.maxShort && gcd(limit, 2*half-4*t) == 1 {
				s.digits = append(s.digits, byte('0'+d))
				s.shorts = t
				s.found = true
				return
			}
		}
		return
	}
	if extra := half - len(s.path) - (limit - s.run); extra > 0 && s.shorts+(extra+limit-1)/limit > s.maxShort {
		return
	}
	for _, d := range []int{0, 2} {
		if s.shorts+d/2 > s.maxShort {
			continue
		}
		next := s.target(cur, d)
		image := s.mirror[next]
		if s.visited[s.sp.class[next]] || s.visited[s.sp.class[image]] || s.sp.class[next] == s.sp.class[image] {
			continue
		}
		mark := len(s.undo)
		for _, q := range []int{cur, s.mirror[cur]} {
			s.kill(q, 2-d)
			s.kill(q, d)
		}
		s.claim(next)
		s.claim(image)
		s.path = append(s.path, next)
		s.digits = append(s.digits, byte('0'+d))
		run := s.run
		if d == 2 {
			s.shorts++
			s.run = 1
		} else {
			s.run++
		}
		if s.feasible() {
			s.dfs()
		}
		if s.found {
			return
		}
		s.shorts -= d / 2
		s.run = run
		s.path = s.path[:len(s.path)-1]
		s.digits = s.digits[:len(s.digits)-1]
		s.visited[s.sp.class[next]] = false
		s.visited[s.sp.class[image]] = false
		s.restore(mark)
	}
}
