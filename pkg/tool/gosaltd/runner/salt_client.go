package runner

import "github.com/funtimecoding/soil/pkg/provision/salt"

func (r *Runner) SaltClient() *salt.Client {
	return r.salt
}
