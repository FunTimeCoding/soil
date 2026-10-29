package runner

import "github.com/funtimecoding/soil/pkg/provision/salt"

func (r *Runner) keep(c *salt.Client) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.closed {
		c.Close()

		return
	}

	r.built = c
}
