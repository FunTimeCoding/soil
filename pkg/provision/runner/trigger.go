package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
)

func (r *Runner) Trigger(request trigger.Request) error {
	select {
	case r.trigger <- request:
		return nil
	default:
		return conflict.Format("run already queued")
	}
}
