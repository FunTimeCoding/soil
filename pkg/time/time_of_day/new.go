package time_of_day

import "time"

func New(
	hour int,
	minute int,
	location *time.Location,
) *Time {
	return &Time{hour: hour, minute: minute, location: location}
}
