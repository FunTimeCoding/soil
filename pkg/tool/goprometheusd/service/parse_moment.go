package service

import "time"

func parseMoment(
	raw string,
	now time.Time,
) (time.Time, error) {
	d, e := time.ParseDuration(raw)

	if e == nil {
		return now.Add(-d), nil
	}

	return time.Parse(time.RFC3339, raw)
}
