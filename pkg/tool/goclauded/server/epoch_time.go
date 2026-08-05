package server

import "time"

func epochTime(v *int64) time.Time {
	if v == nil {
		return time.Time{}
	}

	return time.Unix(*v, 0)
}
