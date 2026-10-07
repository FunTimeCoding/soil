package runner

import "github.com/funtimecoding/soil/pkg/provision/types/update"

func (r *Runner) Sync() (*update.Result, error) {
	return r.provision.Sync()
}
