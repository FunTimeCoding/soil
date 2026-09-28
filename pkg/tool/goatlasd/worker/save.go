package worker

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"time"
)

func (w *Worker) save(
	source string,
	v []*placement.Placement,
) {
	for _, p := range v {
		errors.PanicOnError(w.store.SavePlacement(p))
	}

	w.metric.SetPlacements(source, len(v))
	swept, e := w.store.SweepPlacements(source, time.Now().Add(-w.retention))
	errors.PanicOnError(e)
	w.metric.CountSweptPlacements(source, swept)
}
