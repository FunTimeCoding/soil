package host

import "time"

type Host struct {
	Name   string
	Down   bool
	Since  time.Time
	Reason string
}
