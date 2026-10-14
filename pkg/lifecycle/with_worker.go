package lifecycle

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle/worker"
)

func (l *Lifecycle) WithWorker(w face.Worker) *Lifecycle {
	v := worker.New(w)
	l.component = append(l.component, v)

	if _, okay := w.(face.Drainer); okay {
		l.drainer = append(l.drainer, v)
	}

	return l
}
