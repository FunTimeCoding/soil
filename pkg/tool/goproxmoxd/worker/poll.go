package worker

func (w *Worker) Poll() {
	for _, v := range w.refresh() {
		w.collector.Clear(v.Hypervisor)

		if v.Failure != nil {
			w.logFailure(v.Hypervisor, v.Failure)
			w.collector.SetScrape(v.Hypervisor, false, v.Duration)

			continue
		}

		w.gauge(v)
		w.collector.SetScrape(v.Hypervisor, true, v.Duration)
	}
}
