package service

import "time"

func fableReset(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}

	return &t
}
