package note_writer

func (w *Writers) Writer(note int64) string {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	return w.callsign[note]
}
