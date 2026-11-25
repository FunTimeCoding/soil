package note_writer

import "time"

func (w *Writers) Remember(
	note int64,
	callsign string,
	at time.Time,
) {
	if callsign == "" {
		return
	}

	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.callsign[note] = callsign
	w.written[note] = at
}
