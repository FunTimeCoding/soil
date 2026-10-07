package runner

import "github.com/funtimecoding/soil/pkg/provision/types/trigger"

func (r *Runner) Trigger(request trigger.Request) error {
	return r.provision.Trigger(request)
}
