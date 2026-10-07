package runner

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
	"github.com/funtimecoding/soil/pkg/provision/types/update"
)

func (r *Runner) drainChannels() {
	for {
		select {
		case request := <-r.sync:
			u := update.NewResult()
			u.Error = fmt.Errorf("runner stopped")
			request.Response <- u
		case request := <-r.trigger:
			if request.Response != nil {
				t := trigger.NewResult()
				t.Error = fmt.Errorf("runner stopped")
				request.Response <- t
			}
		default:
			return
		}
	}
}
