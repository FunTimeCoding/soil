package outage_tracker

import "github.com/funtimecoding/soil/pkg/peer/outage_tracker/host"

func (r *Tracker) Down(name string) (bool, *host.Host) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	h, present := r.hosts[name]

	if !present || !h.Down {
		return false, nil
	}

	copied := *h

	return true, &copied
}
