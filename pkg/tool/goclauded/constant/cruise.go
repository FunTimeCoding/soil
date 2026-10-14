package constant

import "time"

type Cruise string

const (
	CruiseOff   Cruise = "off"
	CruiseOn    Cruise = "on"
	CruisePaced Cruise = "paced"
)

const (
	CruiseInterval = 30 * time.Second
	WakeWindow     = time.Hour
)

const (
	CruiseTool = "cruise"
	Mode       = "mode"
	Pace       = "pace"
)

var CruisePace = []time.Duration{
	5 * time.Minute,
	15 * time.Minute,
	30 * time.Minute,
}
