package window

import "time"

func (w *Window) Contains(t time.Time) bool {
	if w.start != nil && t.Before(*w.start) {
		return false
	}

	if w.end != nil && t.After(*w.end) {
		return false
	}

	return true
}
