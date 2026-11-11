package worker

import "context"

func (w *Worker) tick(x context.Context) []error {
	latest, e := w.latest(x)

	if e != nil {
		return []error{e}
	}

	w.status.Reset()
	w.completion.Reset()
	w.size.Reset()
	w.repository.Reset()
	w.damage.Reset()

	for _, s := range latest {
		w.status.WithLabelValues(s.Stanza).Set(float64(s.Status))

		for kind, b := range s.Latest {
			w.completion.WithLabelValues(s.Stanza, kind).Set(
				float64(b.Timestamp.Stop),
			)
			w.size.WithLabelValues(s.Stanza, kind).Set(float64(b.Size.Database))
			w.repository.WithLabelValues(s.Stanza, kind).Set(
				float64(b.Size.Repository.Size),
			)
			w.damage.WithLabelValues(s.Stanza, kind).Set(flag(b.Error))
		}
	}

	return nil
}
