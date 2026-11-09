package worker

import "time"

func (w *Worker) tick() []error {
	start := time.Now()
	w.metrics.LastPollTime.SetToCurrentTime()
	defer func() {
		w.metrics.PollDuration.Observe(time.Since(start).Seconds())
	}()
	o, e := w.record()

	if e != nil {
		return []error{e}
	}

	if o.Pruned > 0 {
		w.logger.Structured("pruned records", "count", o.Pruned)
	}

	w.metrics.AlertsTotal.Add(float64(o.Created))
	w.metrics.AlertsFiring.Set(float64(o.Firing))
	w.metrics.RecordsTotal.Set(float64(o.Total))

	return nil
}
