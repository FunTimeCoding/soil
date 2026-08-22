package runner

import "github.com/funtimecoding/soil/pkg/errors/conflict"

func (r *Runner) Trigger(request TriggerRequest) error {
	select {
	case r.trigger <- request:
		return nil
	default:
		return conflict.Format("run already queued")
	}
}
