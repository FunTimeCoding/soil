package worker

import "time"

func (w *Worker) sleep() bool {
	select {
	case <-w.context.Done():
		return false
	case <-time.After(w.Delay(time.Now())):
		return true
	}
}
