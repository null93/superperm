package alpha

import (
	"errors"
)

var ErrAlphabetEmpty = errors.New("alphabet must be at least 1 character long")
var ErrAlphabetCharacter = errors.New("alphabet must only contain printable ascii characters")
var ErrAlphabetDuplicate = errors.New("alphabet must not contain duplicate characters")
var ErrAlphabetSize = errors.New("alphabets must be the same size")
var ErrCannotTranslate = errors.New("input string has characters not in alphabet")

func Validate(alphabet string) error {
	if len(alphabet) < 1 {
		return ErrAlphabetEmpty
	}
	seen := map[rune]bool{}
	for _, r := range alphabet {
		if r < '!' || r > '~' {
			return ErrAlphabetCharacter
		}
		if seen[r] {
			return ErrAlphabetDuplicate
		}
		seen[r] = true
	}
	return nil
}

func Translate(input, from, to string) (string, error) {
	if len([]rune(from)) != len([]rune(to)) {
		return "", ErrAlphabetSize
	}
	mapping := map[rune]rune{}
	for i, r := range []rune(from) {
		mapping[r] = []rune(to)[i]
	}
	result := []rune{}
	for _, r := range input {
		if _, ok := mapping[r]; !ok {
			return "", ErrCannotTranslate
		}
		result = append(result, mapping[r])
	}
	return string(result), nil
}
