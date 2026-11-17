package outage_tracker

import "github.com/funtimecoding/soil/pkg/peer/outage_tracker/edge"

func (r *Tracker) Observe(
	name string,
	e error,
) *edge.Edge {
	result, refused := r.Triage(name, e)

	if refused != nil {
		return nil
	}

	return result
}
