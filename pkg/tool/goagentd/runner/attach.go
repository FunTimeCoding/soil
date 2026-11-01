package runner

import "github.com/funtimecoding/soil/pkg/face"

func (r *Runner) attach(s face.Stream) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.current = s

	if r.interrupted {
		r.kill()
	}
}
