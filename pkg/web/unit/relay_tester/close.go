package relay_tester

import "github.com/funtimecoding/soil/pkg/errors"

func (r *Relay) Close() {
	errors.LogClose(r.listener)
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, c := range r.open {
		errors.LogClose(c)
	}
}
