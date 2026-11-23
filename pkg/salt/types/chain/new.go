package chain

import "github.com/funtimecoding/soil/pkg/salt/types/top_entry"

func New(
	e *top_entry.Entry,
	states []string,
) *Chain {
	return &Chain{
		Environment: e.Environment,
		Target:      e.Target,
		Live:        e.Live,
		States:      states,
	}
}
