package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
)

func (w *Worker) Collect(
	q context.Context,
	s *gazetteer.Gazetteer,
) {
	for _, c := range w.collectors {
		w.attempt(
			c.Source(),
			func() {
				v, e := c.Collect(q, s)
				errors.PanicOnError(e)
				w.save(c.Source(), v)
			},
		)
	}

	if w.lease == nil {
		return
	}

	w.attempt(constant.SourceLease, func() { w.pollLeases(q, s) })
}
