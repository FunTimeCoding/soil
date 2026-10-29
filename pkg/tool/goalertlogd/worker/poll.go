package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"time"
)

func (w *Worker) Poll() {
	defer w.notify()
	start := time.Now()
	w.lastPoll.Store(start)

	if w.metrics != nil {
		w.metrics.LastPollTime.SetToCurrentTime()
	}

	defer func() {
		if w.metrics != nil {
			w.metrics.PollDuration.Observe(time.Since(start).Seconds())
		}
	}()
	o, e := w.service.Record(start)

	if e != nil {
		if classify.Reportable(e) {
			w.reporter.CaptureException(e)
		} else {
			w.logger.Plain(constant.RecordDeferred, e)
		}

		return
	}

	if o.Pruned > 0 {
		w.logger.Structured("pruned records", "count", o.Pruned)
	}

	if w.metrics != nil {
		w.metrics.AlertsTotal.Add(float64(o.Created))
		w.metrics.AlertsFiring.Set(float64(o.Firing))
		w.metrics.RecordsTotal.Set(float64(o.Total))
	}
}
