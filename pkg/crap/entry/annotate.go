package entry

import (
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/mutation"
	"github.com/funtimecoding/soil/pkg/crap/score"
)

func (e *Entry) Annotate(mutants []*mutation.Mutant) {
	for _, m := range mutants {
		switch m.Status {
		case constant.MutantKilled:
			e.Killed++
		case constant.MutantLived:
			e.Lived = append(e.Lived, m)
		}
	}

	if len(e.Lived) > 0 {
		e.Untrusted = true
		e.Score = score.Score(e.Function.Complexity, 0)
	}
}
