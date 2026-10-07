package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/provision/types/update"
)

func (r *Runner) Sync() (*update.Result, error) {
	request := update.Request{Response: make(chan *update.Result, 1)}

	select {
	case r.sync <- request:
		return <-request.Response, nil
	default:
		return nil, conflict.Format("sync already queued")
	}
}
