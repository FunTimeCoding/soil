package worker

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
	"time"
)

func (w *Worker) saveSightings(
	source string,
	v []*sighting.Sighting,
) {
	for _, e := range v {
		errors.PanicOnError(w.store.SaveSighting(e))
	}

	w.metric.SetSightings(source, len(v))
	swept, e := w.store.SweepSightings(source, time.Now().Add(-w.retention))
	errors.PanicOnError(e)
	w.metric.CountSweptSightings(source, swept)
}
