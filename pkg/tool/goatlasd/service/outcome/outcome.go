package outcome

import "time"

type Outcome struct {
	Source    string
	Sightings bool
	Count     int
	Swept     int64
	Took      time.Duration
	Error     error
}
