package server

import "time"

func epochTime(v *int64) time.Time {
	result := epochTimePointer(v)

	if result == nil {
		return time.Time{}
	}

	return *result
}
