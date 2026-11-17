package outage_tracker

import "github.com/funtimecoding/soil/pkg/peer/outage_tracker/edge"

func (r *Tracker) Handle(
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
