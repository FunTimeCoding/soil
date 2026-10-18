package worker

import "github.com/funtimecoding/soil/pkg/reacher/edge"

func (w *Worker) logFailure(
	hypervisor string,
	e error,
) {
	w.reacher.Handle(
		hypervisor,
		e,
		func(x *edge.Edge) { w.log.Plain("poll hypervisor %s", x) },
		func(f error) {
			w.log.Plain("poll hypervisor %s failed: %v", hypervisor, f)
		},
	)
}
