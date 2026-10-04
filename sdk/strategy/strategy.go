package strategy

import (
	"errors"
	"sort"

	"github.com/null93/superperm/sdk/strategy/egan"
	"github.com/null93/superperm/sdk/strategy/standard"
)

var ErrUnknownStrategy = errors.New("unknown strategy")

type Func func(alphabet string) string

var strategies = map[string]Func{
	"egan":     egan.Generate,
	"standard": standard.Generate,
}

func Get(name string) (Func, error) {
	generate, ok := strategies[name]
	if !ok {
		return nil, ErrUnknownStrategy
	}
	return generate, nil
}

func Names() []string {
	names := []string{}
	for name := range strategies {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
