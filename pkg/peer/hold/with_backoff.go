package hold

import "time"

func (h *Hold) WithBackoff(
	minimum time.Duration,
	maximum time.Duration,
) *Hold {
	h.minimum = minimum
	h.maximum = maximum

	return h
}
