package web

import "time"

func traceClock(stamp string) string {
	t, e := time.Parse(time.RFC3339Nano, stamp)

	if e != nil {
		return stamp
	}

	return t.Local().Format(time.TimeOnly)
}
