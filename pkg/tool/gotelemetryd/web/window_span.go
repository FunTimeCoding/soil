package web

import "time"

func windowSpan(window string) time.Duration {
	d, e := time.ParseDuration(window)

	if e != nil {
		return 24 * time.Hour
	}

	return d
}
