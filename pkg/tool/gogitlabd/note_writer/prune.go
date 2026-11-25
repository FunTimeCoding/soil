package note_writer

import "time"

func (w *Writers) Prune(cutoff time.Time) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	for note, at := range w.written {
		if at.Before(cutoff) {
			delete(w.callsign, note)
			delete(w.written, note)
		}
	}
}
