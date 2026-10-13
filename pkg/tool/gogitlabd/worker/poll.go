package worker

import "slices"

func (w *Worker) Poll() {
	changed := false

	for _, i := range w.service.Inventory().Instances {
		entries, e := w.pollInstance(i.Name)

		if e != nil {
			w.reporter.CaptureException(e)

			continue
		}

		requests, f := w.fetchRequests(i.Name)

		if f != nil {
			w.reporter.CaptureException(f)

			continue
		}

		w.mutex.Lock()

		if !slices.EqualFunc(w.entries[i.Name], entries, sameEntry) ||
			!slices.EqualFunc(w.requests[i.Name], requests, sameRequest) {
			changed = true
		}

		w.entries[i.Name] = entries
		w.requests[i.Name] = requests
		w.mutex.Unlock()
	}

	w.refreshGauge()

	if changed {
		w.notifier.Notify()
	}
}
