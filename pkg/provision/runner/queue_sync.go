package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/types/update"
)

func (r *Runner) queueSync(request update.Request) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.closed {
		return unavailable.Format(constant.RunnerStopped)
	}

	select {
	case r.sync <- request:
		return nil
	default:
		return conflict.Format("sync already queued")
	}
}
