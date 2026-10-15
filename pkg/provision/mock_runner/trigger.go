package mock_runner

import "github.com/funtimecoding/soil/pkg/provision/types/trigger"

func (r *Runner) Trigger(request trigger.Request) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.fail != nil {
		return r.fail
	}

	r.requests = append(r.requests, request)

	return nil
}
