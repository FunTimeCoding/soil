package entry

import (
	"github.com/funtimecoding/soil/pkg/crap/function"
	"github.com/funtimecoding/soil/pkg/crap/score"
)

func New(
	f *function.Function,
	coverage float64,
) *Entry {
	return &Entry{
		Function: f,
		Coverage: coverage,
		Score:    score.Score(f.Complexity, coverage),
	}
}
