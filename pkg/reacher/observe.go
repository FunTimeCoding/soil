package reacher

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/reacher/edge"
)

func (r *Reacher) Observe(
	name string,
	e error,
) *edge.Edge {
	if e == nil {
		return r.Succeed(name)
	}

	f := connection.Classify(e)

	if f == nil {
		return nil
	}

	return r.Fail(name, f.Reason)
}
