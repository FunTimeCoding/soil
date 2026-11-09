package window

import "time"

func (w *Window) Bound(
	now time.Time,
	span time.Duration,
) *Window {
	if w.end == nil {
		w.end = &now
	}

	if w.start == nil {
		start := w.end.Add(-span)
		w.start = &start
	}

	return w
}
