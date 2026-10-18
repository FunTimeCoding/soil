package edge

import "time"

type Edge struct {
	Host     string
	Down     bool
	Reason   string
	Duration time.Duration
}
