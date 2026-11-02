package runner

import "github.com/funtimecoding/soil/pkg/provision/types/update"

func (r *Runner) Sync() (*update.Result, error) {
	request := update.Request{Response: make(chan *update.Result, 1)}

	if e := r.queueSync(request); e != nil {
		return nil, e
	}

	return <-request.Response, nil
}
