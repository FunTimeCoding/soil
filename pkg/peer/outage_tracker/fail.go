package outage_tracker

import (
	"github.com/funtimecoding/soil/pkg/peer/outage_tracker/edge"
	"github.com/funtimecoding/soil/pkg/peer/outage_tracker/host"
)

func (r *Tracker) Fail(
	name string,
	reason string,
) *edge.Edge {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	h, present := r.hosts[name]

	if !present {
		h = host.New(name)
		r.hosts[name] = h
	}

	if h.Down {
		return nil
	}

	h.Down = true
	h.Since = r.clock()
	h.Reason = reason

	return edge.Down(name, reason)
}
