package worker

import "github.com/funtimecoding/soil/pkg/face"

func New(w face.Worker) *Worker {
	return &Worker{Face: w}
}
