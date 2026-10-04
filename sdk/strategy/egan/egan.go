package egan

import (
	"github.com/null93/superperm/sdk/strategy/standard"
)

func Generate(alphabet string) string {
	n := len(alphabet)
	if n < 4 {
		return standard.Generate(alphabet)
	}
	fact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * i
	}
	out := make([]uint8, 0, fact[n]+fact[n-1]+fact[n-2]+fact[n-3]+n-3)
	for d := n - 1; d >= 2; d-- {
		out = append(out, uint8(d))
	}
	out = append(out, uint8(n), 1)
	q := []uint8{1, uint8(n)}
	for d := n - 1; d >= 2; d-- {
		q = append(q, uint8(d))
	}
	for visited := 1; visited < fact[n]; visited++ {
		c := out[len(out)-n:]
		if weight2(c) && !equal(c, q) {
			out = append(out, c[1], c[0])
		} else {
			out = append(out, c[0])
		}
	}
	mapping := make([]byte, n+1)
	for i := 0; i < n; i++ {
		mapping[out[i]] = alphabet[i]
	}
	result := make([]byte, len(out))
	for i, d := range out {
		result[i] = mapping[d]
	}
	return string(result)
}

func weight2(c []uint8) bool {
	n := len(c)
	f := int(c[0])
	if f == n {
		return false
	}
	pos := 1
	for int(c[pos]) != n {
		pos++
	}
	want := f - 1
	if want == 0 {
		want = n - 1
	}
	return int(c[1+pos%(n-1)]) == want
}

func equal(a, b []uint8) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
