package worker

import "time"

func (w *Worker) Poll() {
	for _, i := range w.service.Instances() {
		start := time.Now()
		e := w.pollInstance(i.Name)

		if e != nil {
			w.log.Plain("poll hypervisor %s failed: %v", i.Name, e)
			w.collector.Clear(i.Name)
			w.collector.SetScrape(i.Name, false, time.Since(start))

			continue
		}

		w.collector.SetScrape(i.Name, true, time.Since(start))
	}
}
