package watch

func (w *Watch) tick() []error {
	w.Poll()

	return nil
}
