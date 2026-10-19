package reacher

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/peer/reacher/edge"
)

func (r *Reacher) Triage(
	name string,
	e error,
) (*edge.Edge, error) {
	if e == nil {
		return r.Succeed(name), nil
	}

	f := connection.Classify(e)

	if f == nil {
		return nil, e
	}

	return r.Outage(name, f.Reason, e), nil
}
