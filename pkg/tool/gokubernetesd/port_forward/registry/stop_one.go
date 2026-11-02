package registry

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (r *Registry) StopOne(identifier string) error {
	r.mutex.Lock()
	f, okay := r.forwards[identifier]
	delete(r.forwards, identifier)
	r.mutex.Unlock()

	if !okay {
		return not_found.New("port forward", identifier)
	}

	f.Stop()

	return nil
}
