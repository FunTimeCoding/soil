package window

import "time"

func (w *Window) Start() (time.Time, bool) {
	if w.start == nil {
		return time.Time{}, false
	}

	return *w.start, true
}
