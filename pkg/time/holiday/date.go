package holiday

import (
	"github.com/funtimecoding/soil/pkg/time/day"
	"time"
)

func (h *Holiday) Date(year int) time.Time {
	if h.movable {
		return Easter(year).AddDate(0, 0, h.offset)
	}

	return day.New(year, h.month, h.day)
}
