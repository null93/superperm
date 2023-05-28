package problem

import (
	"errors"
	"time"

	"github.com/null93/superperm/sdk/problem/rotate"
)

var ErrUnsupportedStrategy = errors.New("unsupported strategy")

var SupportedStrategy = []string{"rotate"}

type Solvable interface {
	Solve(string)
	GetSolution() string
	GetLength() int
	GetDuration() time.Duration
}

func IsSupportedStrategy(strategy string) bool {
	for _, s := range SupportedStrategy {
		if s == strategy {
			return true
		}
	}
	return false
}

func NewProblem(strategy string) (Solvable, error) {
	switch strategy {
	case "rotate":
		return &rotate.Problem{}, nil
	}
	return nil, ErrUnsupportedStrategy
}
