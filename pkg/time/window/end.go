package window

import "time"

func (w *Window) End() (time.Time, bool) {
	if w.end == nil {
		return time.Time{}, false
	}

	return *w.end, true
}
