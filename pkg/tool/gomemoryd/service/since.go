package service

import "time"

func since(
	now time.Time,
	window time.Duration,
) string {
	return now.Add(-window).UTC().Format(time.RFC3339)
}
