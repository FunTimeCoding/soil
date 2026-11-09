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

const (
	CruisePaceMinimumMinutes = 1
	CruisePaceMaximumMinutes = 60
)
