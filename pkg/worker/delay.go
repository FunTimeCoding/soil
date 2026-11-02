package worker

import "time"

func (w *Worker) Delay(now time.Time) time.Duration {
	if w.at == nil {
		return w.interval
	}

	return w.at.Next(now).Sub(now)
}
