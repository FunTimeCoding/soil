package reacher

import "github.com/funtimecoding/soil/pkg/reacher/host"

func (r *Reacher) Down(name string) (bool, *host.Host) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	h, present := r.hosts[name]

	if !present || !h.Down {
		return false, nil
	}

	copied := *h

	return true, &copied
}
