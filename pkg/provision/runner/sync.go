package runner

import "github.com/funtimecoding/soil/pkg/errors/conflict"

func (r *Runner) Sync() (*SyncResult, error) {
	request := SyncRequest{Response: make(chan *SyncResult, 1)}

	select {
	case r.sync <- request:
		return <-request.Response, nil
	default:
		return nil, conflict.Format("sync already queued")
	}
}
