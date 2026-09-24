package entry

import (
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/function"
)

func NewMissing(
	f *function.Function,
	p constant.Policy,
) *Entry {
	coverage := 0.0

	if p == constant.Optimistic {
		coverage = 100
	}

	result := New(f, coverage)
	result.Missing = true

	return result
}
