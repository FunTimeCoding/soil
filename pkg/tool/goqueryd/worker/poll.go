package worker

import "github.com/funtimecoding/soil/pkg/errors"

func (w *Worker) poll() {
	_, e := w.service.IndexCollections("")
	errors.PanicOnError(e)
	w.service.EmbedPending(w.context)
}
