package reacher

import "github.com/funtimecoding/soil/pkg/reacher/edge"

func (r *Reacher) Succeed(name string) *edge.Edge {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	h, present := r.hosts[name]

	if !present || !h.Down {
		return nil
	}

	h.Down = false
	h.Reason = ""

	return edge.Up(name, r.clock().Sub(h.Since))
}
