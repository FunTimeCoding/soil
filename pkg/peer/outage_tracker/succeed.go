package outage_tracker

import "github.com/funtimecoding/soil/pkg/peer/outage_tracker/edge"

func (r *Tracker) Succeed(name string) *edge.Edge {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	h, present := r.hosts[name]

	if !present || !h.Down {
		return nil
	}

	h.Down = false
	h.Reason = ""
	h.Reported = false

	return edge.Up(name, r.clock().Sub(h.Since))
}
