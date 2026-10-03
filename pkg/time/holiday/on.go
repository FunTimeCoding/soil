package holiday

import "time"

func (h *Holiday) On(t time.Time) bool {
	d := h.Date(t.Year())

	return d.Month() == t.Month() && d.Day() == t.Day()
}
