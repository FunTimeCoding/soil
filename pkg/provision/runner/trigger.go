package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
)

func (r *Runner) Trigger(request trigger.Request) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.closed {
		return unavailable.Format(constant.RunnerStopped)
	}

	select {
	case r.trigger <- request:
		return nil
	default:
		return conflict.Format("run already queued")
	}
}
