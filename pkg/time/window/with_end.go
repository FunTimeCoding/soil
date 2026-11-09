package window

import "time"

func (w *Window) WithEnd(t time.Time) *Window {
	w.end = &t

	return w
}
