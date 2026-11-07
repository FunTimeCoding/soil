package lifecycle

import "github.com/funtimecoding/soil/pkg/face"

func (l *Lifecycle) WithWorkers(workers ...face.Worker) *Lifecycle {
	for _, w := range workers {
		l.WithWorker(w)
	}

	return l
}
