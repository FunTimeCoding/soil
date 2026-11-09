package window

import "time"

func (w *Window) WithStart(t time.Time) *Window {
	w.start = &t

	return w
}
