package reacher

import (
	"github.com/funtimecoding/soil/pkg/peer/constant"
	"github.com/funtimecoding/soil/pkg/peer/outage"
)

func (r *Reacher) lasting(
	name string,
	e error,
) *outage.OutageError {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	h, present := r.hosts[name]

	if !present || !h.Down || h.Reported {
		return nil
	}

	took := r.clock().Sub(h.Since)

	if took < constant.Lasting {
		return nil
	}

	h.Reported = true

	return outage.New(name, h.Since, took, e)
}
