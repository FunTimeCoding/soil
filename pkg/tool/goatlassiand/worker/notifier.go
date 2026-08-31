package worker

import "github.com/funtimecoding/soil/pkg/face"

func (w *Worker) Notifier() face.EventNotifier {
	return w.notifier
}
