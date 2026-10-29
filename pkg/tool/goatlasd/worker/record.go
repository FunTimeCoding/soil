package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/service/outcome"
)

func (w *Worker) record(o *outcome.Outcome) {
	w.metric.ObserveCollect(o.Source, o.Took)
	w.metric.CountFailure(o.Source, o.Error != nil)

	switch {
	case unavailable.Is(o.Error):
		w.logger.Plain(constant.SourceAwayFormat, o.Source, o.Error)
	case o.Error != nil:
		w.reporter.CaptureException(o.Error)
	case o.Sightings:
		w.metric.SetSightings(o.Source, o.Count)
		w.metric.CountSweptSightings(o.Source, o.Swept)
	default:
		w.metric.SetPlacements(o.Source, o.Count)
		w.metric.CountSweptPlacements(o.Source, o.Swept)
	}
}
