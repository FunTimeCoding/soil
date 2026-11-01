package service

import "time"

func parseTimestamp(ts string) time.Time {
	t, e := time.Parse(time.RFC3339Nano, ts)

	if e != nil {
		t, e = time.Parse("2006-01-02T15:04:05.000Z", ts)
	}

	if e != nil {
		return time.Now()
	}

	return t.Local()
}
