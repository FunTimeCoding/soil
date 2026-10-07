package runner_tester

import "github.com/funtimecoding/soil/pkg/provision/types/apply_call"

func (o *Tester) LastApply() *apply_call.Call {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.applied[len(o.applied)-1]
}
