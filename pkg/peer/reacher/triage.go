package reacher

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/peer/reacher/edge"
)

func (r *Reacher) Triage(
	name string,
	e error,
) (*edge.Edge, error) {
	if e == nil {
		return r.Succeed(name), nil
	}

	if f := connection.Classify(e); f != nil {
		return r.Outage(name, f.Reason, e), nil
	}

	if unavailable.Is(e) {
		return r.Outage(name, e.Error(), e), nil
	}

	return nil, e
}
