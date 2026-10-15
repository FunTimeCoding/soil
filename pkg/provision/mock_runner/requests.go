package mock_runner

import "github.com/funtimecoding/soil/pkg/provision/types/trigger"

func (r *Runner) Requests() []trigger.Request {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return append([]trigger.Request{}, r.requests...)
}
