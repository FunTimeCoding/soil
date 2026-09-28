package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
)

func (w *Worker) pollLeases(
	q context.Context,
	s *gazetteer.Gazetteer,
) {
	v, e := w.lease.Collect(q, s)
	errors.PanicOnError(e)
	w.saveSightings(constant.SourceLease, v)
}
