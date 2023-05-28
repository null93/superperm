package alpha

import (
	"errors"
)

var ErrUnknownAlphabet = errors.New("unknown alphabet")
var ErrAlphabetSize = errors.New("cannot use alphabet with given n")
var ErrCannotTranslate = errors.New("input string has characters not in alphabet")

var Alphabets = map[string]string{
	"latin": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"digit": "1234567890",
	"hex":   "0123456789ABCDEF",
}

type AlphabetMapping struct {
	From    string
	To      string
	Mapping map[rune]rune
}

type Translator interface {
	Translate(input string) (string, error)
}

func IsSupportedAlphabet(alphabet string) bool {
	_, ok := Alphabets[alphabet]
	return ok
}

func (m *AlphabetMapping) Translate(input string) (string, error) {
	result := []rune{}
	for _, r := range input {
		if _, ok := m.Mapping[r]; !ok {
			return "", ErrCannotTranslate
		}
		result = append(result, m.Mapping[r])
	}
	return string(result), nil
}

func NewTranslator(from, to string, n int) (Translator, error) {
	if _, ok := Alphabets[from]; !ok {
		return nil, ErrUnknownAlphabet
	}
	if _, ok := Alphabets[to]; !ok {
		return nil, ErrUnknownAlphabet
	}
	if len(Alphabets[from]) < n || len(Alphabets[to]) < n {
		return nil, ErrAlphabetSize
	}
	am := &AlphabetMapping{
		From:    Alphabets[from][0:n],
		To:      Alphabets[to][0:n],
		Mapping: map[rune]rune{},
	}
	for i, r := range am.From {
		am.Mapping[r] = []rune(am.To)[i]
	}
	return am, nil
}

func GetAlphabet(name string, n int) (string, error) {
	if _, ok := Alphabets[name]; !ok {
		return "", ErrUnknownAlphabet
	}
	if len(Alphabets[name]) < n {
		return "", ErrAlphabetSize
	}
	return Alphabets[name][0:n], nil
}
