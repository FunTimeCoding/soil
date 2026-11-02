package time_of_day

import "time"

func (t *Time) Next(now time.Time) time.Time {
	local := now.In(t.location)
	y, m, d := local.Date()
	result := time.Date(y, m, d, t.hour, t.minute, 0, 0, t.location)

	if result.After(now) {
		return result
	}

	return time.Date(y, m, d+1, t.hour, t.minute, 0, 0, t.location)
}
