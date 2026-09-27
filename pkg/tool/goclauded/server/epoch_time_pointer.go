package server

import "time"

func epochTimePointer(v *int64) *time.Time {
	if v == nil {
		return nil
	}

	return new(time.Unix(*v, 0).UTC())
}
