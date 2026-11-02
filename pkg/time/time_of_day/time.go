package time_of_day

import "time"

type Time struct {
	hour     int
	minute   int
	location *time.Location
}
