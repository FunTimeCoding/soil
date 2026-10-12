package watch

func (w *Watch) Poll() {
	s, e := w.source.Health()

	if e != nil {
		w.reachable.Set(0)

		return
	}

	w.reachable.Set(1)

	if s.Sealed {
		w.sealed.Set(1)
	} else {
		w.sealed.Set(0)
	}
}
