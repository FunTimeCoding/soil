package reacher

import "github.com/funtimecoding/soil/pkg/reacher/edge"

func (r *Reacher) Observe(
	name string,
	e error,
) *edge.Edge {
	result, refused := r.Triage(name, e)

	if refused != nil {
		return nil
	}

	return result
}
