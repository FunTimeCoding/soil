package registry

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward"

func (r *Registry) Launch(
	f *port_forward.Forward,
	run func(),
) {
	r.mutex.Lock()
	r.forwards[f.Identifier] = f
	r.mutex.Unlock()
	r.running.Add(1)
	go func() {
		defer r.running.Done()
		run()
		r.mutex.Lock()

		if r.forwards[f.Identifier] == f {
			delete(r.forwards, f.Identifier)
		}

		r.mutex.Unlock()
	}()
}
