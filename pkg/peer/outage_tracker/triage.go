package outage_tracker

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/remediable"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/peer/outage_tracker/edge"
)

func (r *Tracker) Triage(
	name string,
	e error,
) (*edge.Edge, error) {
	if e == nil {
		return r.Succeed(name), nil
	}

	if remediable.Is(e) {
		return nil, e
	}

	if f := connection.Classify(e); f != nil {
		return r.Outage(name, f.Reason, e), nil
	}

	if unavailable.Is(e) {
		return r.Outage(name, e.Error(), e), nil
	}

	return nil, e
}
