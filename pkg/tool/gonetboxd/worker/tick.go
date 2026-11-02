package worker

func (w *Worker) tick() []error {
	v, e := w.service.Intents()

	if e != nil {
		return []error{e}
	}

	w.metric.SetIntents(v)

	return nil
}
