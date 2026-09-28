package worker

import "time"

func (w *Worker) attempt(
	source string,
	f func(),
) {
	started := time.Now()
	completed := false
	w.recovery.Run(
		func() {
			f()
			completed = true
		},
	)
	w.metric.ObserveCollect(source, time.Since(started))
	w.metric.CountFailure(source, !completed)
}
