package reacher

import "github.com/funtimecoding/soil/pkg/peer/reacher/edge"

func (r *Reacher) Handle(
	name string,
	e error,
	crossed func(*edge.Edge),
	refused func(error),
) {
	result, real := r.Triage(name, e)

	if result != nil {
		crossed(result)
	}

	if real == nil || !r.firstRefusal(name, real) {
		return
	}

	refused(real)
}
