package reacher

import (
	"github.com/funtimecoding/soil/pkg/reacher/edge"
	"github.com/funtimecoding/soil/pkg/reacher/host"
)

func (r *Reacher) Fail(
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
