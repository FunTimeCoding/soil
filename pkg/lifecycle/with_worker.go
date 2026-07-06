package lifecycle

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle/worker"
)

func WithWorker(w face.Worker) Option {
	return func(l *Lifecycle) {
		l.component = append(
			l.component,
			worker.New(w),
		)
	}
}
