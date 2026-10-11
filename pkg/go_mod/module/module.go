package module

import "time"

type Module struct {
	MonitorIdentifier string
	Path              string
	Version           string
	Latest            string
	Release           *time.Time
	Months            int
	Indirect          bool
	Untagged          bool
	now               time.Time
	threshold         int
	concern           []string
}
