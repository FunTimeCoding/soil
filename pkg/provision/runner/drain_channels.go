package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
	"github.com/funtimecoding/soil/pkg/provision/types/update"
)

func (r *Runner) drainChannels() {
	for {
		select {
		case request := <-r.sync:
			u := update.NewResult()
			u.Error = unavailable.Format(constant.RunnerStopped)
			request.Response <- u
		case request := <-r.trigger:
			if request.Response != nil {
				t := trigger.NewResult()
				t.Error = unavailable.Format(constant.RunnerStopped)
				request.Response <- t
			}
		default:
			return
		}
	}
}
